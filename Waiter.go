package main

import (
	"fmt"
	"github.com/Anniegavr/Lobby/Lobby/models"
	"github.com/Anniegavr/Lobby/Lobby/utils"

	//"Table.go"
	_ "fmt"
	"time"
	_ "time"
)

type Status int

const (
	Free             Status = 0
	Busy                    = 1
	LookingForOrders        = 2
)

type Waiter struct {
	id int
	assignedTables []*Table
	queue          *models.Queue
	//status chan bool
	conf *Configuration
}

func NewWaiter(id int, conf *Configuration) *Waiter {
	return &Waiter{
		id:    id,
		conf:  conf,
		queue: models.NewQueue(),
	}
}

func (waiter *Waiter) GetId() int {
	return waiter.id
}

func (waiter *Waiter) Run() {
	for {
		waiter.update()
	}
}


func contains(s []*Table, e *Table) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (waiter *Waiter) AddTable(table *Table) {
	if contains(waiter.assignedTables, table){
		return
	}

	waiter.assignedTables = append(waiter.assignedTables, table)
}


func (waiter *Waiter) AddDistributionData(data *utils.DistributionData) {
	waiter.queue.Push(data)
}

func (waiter *Waiter) update() {
	orderDone := waiter.queue.Len() != 0
	if orderDone {
		for waiter.queue.Len() != 0 {
			data := waiter.popData()
			tab := waiter.getTableById(data.TableID)
			tab.GetOrder(data)

			fmt.Println("Time of getting order = ", time.Now())
			fmt.Println(*data)
			fmt.Println()
		}
	}

	for _, tab := range waiter.assignedTables {
		if tab.GetOrderingStatus() {
			err := tab.StartOrdering()
			if err != nil {
				continue
			}

			timeToMakeOrder := Range(waiter.conf.MinMakeOrder, waiter.conf.MaxMakeOrder)
			durationToMakeOrder := time.Duration(timeToMakeOrder)
			//time.Sleep(configuration.TimeUnit * durationToMakeOrder)
			time.Sleep(models.TimeUnit * durationToMakeOrder)

			order, err := tab.FinishOrdering(waiter.id)
			if err != nil {
				continue
			}

			SendOrder(order, waiter.conf)

			fmt.Println("Time of sending order = ", time.Now())
			fmt.Println(*order)
			fmt.Println()
		}
	}
}

func (waiter *Waiter) getTableById(id int) *Table {
	for _, table := range waiter.assignedTables {
		if table.GetId() == id {
			return table
		}
	}
	return nil
}

func (waiter *Waiter) popData() *utils.DistributionData {
	dataRef := waiter.queue.Pop()
	data := dataRef.(*utils.DistributionData)

	return data
}

