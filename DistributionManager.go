package main

import (
	"github.com/Anniegavr/Lobby/Lobby/utils"
)

var waiters []*Waiter

func SetWaiters(ws []*Waiter) {
	waiters = ws
}

func PushQueue(data *utils.DistributionData) {
	for _, w := range waiters {
		if data.WaiterID == w.GetId() {
			w.AddDistributionData(data)
			return
		}
	}
}
