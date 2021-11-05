package main

import (
	"errors"
	"fmt"
	"github.com/Anniegavr/Lobby/Lobby/models"
	"github.com/Anniegavr/Lobby/Lobby/models/item"
	"github.com/Anniegavr/Lobby/Lobby/utils"
	"math/rand"
	"sync"
	"time"
)

type TableStatus int

const (
	FreeForClients TableStatus = 0
	Ordering                   = 1
	Waiting                    = 2
)

type Table struct {
	id     int
	status TableStatus

	orderStatus chan bool

	mutex sync.Mutex
	manager *TableIdCounter
	menu    *item.Container
	conf    *Configuration
	rate *models.RatingSystem
}
func NewTable(
	id int,
	manager *TableIdCounter,
	menu *item.Container,
	conf *Configuration) *Table {
	return &Table{
		id:          id,
		manager:     manager,
		menu:        menu,
		conf:        conf,
		orderStatus: make(chan bool),
	}
}

func (table *Table) GetId() int {
	return table.id
}

func (table *Table) GetStatus() TableStatus {
	return table.status
}

func (table *Table) GetOrderingStatus() bool {
	return table.status == Ordering
}


func (table *Table) Run() {
	for {
		table.update()
	}
}

func (table *Table) StartOrdering() error {
	table.mutex.Lock()

	if table.status != Ordering {
		return errors.New("can't place order")
	}
	table.status = Waiting

	table.mutex.Unlock()

	return nil
}

func (table *Table) FinishOrdering(waiterId int) (*models.OrderData, error) {
	priority := table.getPriority()
	count := table.getOrderCount()

	items := make([]int, count)
	var maxWait int = 0

	for i := 0; i < count; i++ {
		menuLen := table.menu.GetLen()
		index := rand.Intn(menuLen)
		tab, ok := table.menu.Get(index)
		if ok != true {
			return nil, errors.New("outbound array index")
		}

		items[i] = tab.Id
		if maxWait < tab.PreparationTime{
			maxWait = tab.PreparationTime
		}
	}

	finalMaxWait := float32(maxWait) * table.conf.MaxWaitMultiplier
	pickUpTime := time.Now().Unix()

	order := &models.OrderData{
		OrderId:    table.manager.Get(),
		TableId:    table.id,
		WaiterId:   waiterId,
		Items:      items,
		Priority:   priority,
		MaxWait:    finalMaxWait,
		PickUpTime: pickUpTime,
	}

	return order, nil
}


func (table *Table) GetOrder(dist *utils.DistributionData) {
	<-table.orderStatus

	now := time.Now().Unix()
	rating := models.Calculate(dist.PickUpTime, now, dist.MaxWait)

	fmt.Printf("%s = %d\n", "Rating order", rating)

	table.rate.Add(rating)

	fmt.Printf("%s = %f\n", "Rating overall", table.rate.Return())
}

func (table *Table) SetRatingSystem(rate *models.RatingSystem) {
	table.rate = rate
}

func (table *Table) update() {
	table.free()
	table.makingOrder()
}

func (table *Table) free() {
	table.status = FreeForClients
	time.Sleep(models.TimeUnit)
}

func (table *Table) makingOrder() {
	table.status = Ordering
	table.orderStatus <- true
}

func (table *Table) getPriority() int {
	minPriority := table.conf.MinPriority
	maxPriority := table.conf.MaxPriority

	priority := Range(minPriority, maxPriority)

	return priority
}

func (table *Table) getOrderCount() int {
	minOrderItems := table.conf.MinOrderItems
	maxOrderItems := table.conf.MaxOrderItems

	orderItemsCount := Range(minOrderItems, maxOrderItems)

	return orderItemsCount
}
