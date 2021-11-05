package main

import (
	"encoding/json"
	"fmt"
	"github.com/Anniegavr/Lobby/Lobby/models"
)

const (
	HttpAddr = "http://"
)

func SendOrder(order *models.OrderData, conf *Configuration) {
	//orderIdentif++
	//	items := []int{3, 4, 4, 5}
	//	order := Order{
	//		orderIdentif, 1, 1, items, 3, 45, 1631453140,
	//	}
	//
	//	resp, err := http.Post("http://localhost:8082/orders", "application/json",
	//		bytes.NewBuffer(jsonData))
	addr := HttpAddr + conf.KitchenAddr + conf.OrderRout

	jsonBuff, err := json.Marshal(*order)
	if err != nil {
		fmt.Println(err)
		return
	}

	SendRequest(addr, jsonBuff)
}
