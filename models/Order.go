package models


import "fmt"

type OrderData struct {
	OrderId  int `json:"order_id"`
	TableId  int   `json:"table_id"`
	WaiterId int   `json:"waiter_id"`
	Items    []int `json:"items"`
	Priority   int     `json:"priority"`
	MaxWait    float32 `json:"max_wait"`
	PickUpTime int64   `json:"pick_up_time"`
}

func (data OrderData) String() string {
	var result string

	result += fmt.Sprintln("OrderId = ", data.OrderId)
	result += fmt.Sprintln("TableId = ", data.TableId)
	result += fmt.Sprintln("WaiterId = ", data.WaiterId)
	result += fmt.Sprintln("Items = ", data.Items)
	result += fmt.Sprintln("Priority = ", data.Priority)
	result += fmt.Sprintln("MaxWait = ", data.MaxWait)
	result += fmt.Sprintln("PickUpTime = ", data.PickUpTime)

	return result
}
