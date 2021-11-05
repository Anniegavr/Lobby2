package utils
import "fmt"

type DistributionData struct {
	OrderID        int              `json:"order-id"`
	TableID        int              `json:"table-id"`
	WaiterID       int              `json:"waiter-id"`
	Items          []int            `json:"items"`
	Priority       int              `json:"priority"`
	MaxWait        float32          `json:"max-wait"`
	PickUpTime     int64            `json:"pick-up-time"`
	CookingTime    int              `json:"cooking-time"`
	CookingDetails []CookingDetails `json:"cooking-details"`
}

func (data DistributionData) String() string {
	var result string

	result += fmt.Sprintln("OrderId = ", data.OrderID)
	result += fmt.Sprintln("TableId = ", data.TableID)
	result += fmt.Sprintln("WaiterId = ", data.WaiterID)
	result += fmt.Sprintln("Items = ", data.Items)
	result += fmt.Sprintln("Priority = ", data.Priority)
	result += fmt.Sprintln("MaxWait = ", data.MaxWait)
	result += fmt.Sprintln("PickUpTime = ", data.PickUpTime)
	result += fmt.Sprintln("CookingTime = ", data.CookingTime)
	result += fmt.Sprintln("CookingDetails = ", data.CookingDetails)

	return result
}
