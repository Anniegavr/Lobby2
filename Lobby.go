package main

import (
	"encoding/json"
	"github.com/Anniegavr/Lobby/Lobby/models"
	item2 "github.com/Anniegavr/Lobby/Lobby/models/item"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)


//type CookingDetails struct {
//	FoodId int `json:"food_id"`
//	CookId int `json:"cook_id"`
//}

//type Order struct {
//	OrderId  int   `json:"id"`
//	TableId  int   `json:"table_id"`
//	WaiterId int   `json:"table_id"`
//	Items    []int `json:"items"`
//	Priority int   `json:"priority"`
//	MaxWait  int   `json:"maxWait"`
//	PickTime int   `json:"pick_time"`
//}



//type TableIdCounter struct {
//	id int
//}
//func New() *TableIdCounter {
//	return &TableIdCounter{
//		id: 0,
//	}
//}
//
//func (counter *TableIdCounter) getIdRef() *int {
//	return &counter.id
//}
//
//func (counter *TableIdCounter) Get() int {
//	idRef := counter.getIdRef()
//	defer func() {
//		*idRef++
//	}()
//
//	return *idRef
//}

//type Menu []item2.Item
//type OrdersList []models.OrderData
//type Conf struct {
//	LobbyAddr                string  `json:"lobby_addr"`
//	KitchenAddr                   string  `json:"kitchen_addr"`
//	DistributionRout              string  `json:"distribution_rout"`
//	OrderRout                     string  `json:"order_rout"`
//	TableCount                    int     `json:"table_count"`
//	WaiterCount                   int     `json:"waiter_count"`
//	MinMakeOrder                  int     `json:"min_make_order"`
//	MaxMakeOrder                  int     `json:"max_make_order"`
//	MinOrderItems                 int     `json:"min_order_items"`
//	MaxOrderItems                 int     `json:"max_order_items"`
//	MinPriority                   int     `json:"min_priority"`
//	MaxPriority                   int     `json:"max_priority"`
//	TimeUnitMillisecondMultiplier int     `json:"time_unit_millisecond_multiplier"`
//	MaxWaitMultiplier             float64 `json:"max_wait_multiplier"`
//}
const (
	ConfPath  = "./configuration.json"
	ItemsPath = "./items.json"
)

//var sendOrder SendOrder = new
//func generateOrder() {
//	i := 1
//	max := 10
//	for i <= max {
//		// wait for 3-10 seconds betwwen placing orders
//		preparationTime := rand.Intn(5)
//		time.Sleep(time.Duration(preparationTime) * time.Second)
//		//sendOrder.SendOrder()
//		i += 1
//	}
//}

//type orders []int                                      //a list of orders
//func indexPage(w http.ResponseWriter, r *http.Request) {}



//func handleRequests() {
//	myRouter := mux.NewRouter().StrictSlash(true)
//	myRouter.HandleFunc("/", indexPage)
//	myRouter.HandleFunc("/distribution", SendRequest()).Methods("POST")
//	log.Fatal(http.ListenAndServe(":8083", myRouter))
//}

func GetItemContainer() *item2.Container {
	var itemList []item2.Item

	itemListFile, _ := os.Open(ItemsPath)
	defer func(itemListFile *os.File) {
		_ = itemListFile.Close()
	}(itemListFile)

	jsonData, err := io.ReadAll(itemListFile)
	if err != nil {
		log.Fatalf("exit: %s\n", err.Error())
		return nil
	}

	err = json.Unmarshal(jsonData, &itemList)
	if err != nil {
		log.Fatalf("exit: %s\n", err.Error())
		return nil
	}

	return item2.NewContainer(itemList)
}

func main() {
	conf := GetConf()
	container := GetItemContainer()

	rate := models.NewRatingSystem()

	timeUnitMillisecondMultiplier := time.Duration(conf.TimeUnitMillisecondMultiplier)
	models.TimeUnit = time.Millisecond * timeUnitMillisecondMultiplier

	manager := NewTableIdCounter()

	var tables = make([]*Table, conf.TableCount)
	for index := range tables {
		tables[index] = NewTable(index, manager, container, &conf)
	}

	var waiters = make([]*Waiter, conf.WaiterCount)
	for index := range waiters {
		waiters[index] = NewWaiter(index, &conf)
	}

	SetWaiters(waiters)

	for i, e := range tables {
		waiters[i%conf.WaiterCount].AddTable(e)
	}

	for index := range tables {
		tables[index].SetRatingSystem(rate)
		go tables[index].Run()
	}

	for index := range waiters {
		go waiters[index].Run()
	}

	http.HandleFunc(conf.DistributionRout, DistributionHandler)

	err := http.ListenAndServe(conf.LobbyAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConf() Configuration {
	var conf Configuration

	confFile, _ := os.Open(ConfPath)
	defer func(confFile *os.File) {
		_ = confFile.Close()
	}(confFile)

	jsonData, err := io.ReadAll(confFile)
	if err != nil {
		log.Fatalf("exit: %s\n", err.Error())
		return conf
	}

	err = json.Unmarshal(jsonData, &conf)
	if err != nil {
		log.Fatalf("exit: %s\n", err.Error())
		return conf
	}

	return conf
}
