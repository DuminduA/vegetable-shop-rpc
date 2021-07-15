package main

import (
	"fmt"
	"net/rpc"
	common "vegetable-shop-rpc/client/model"
)

func main() {
	var veg common.Vegetable
	var allVeg []string
	var availableAmount float32
	var price float32
	var priceDto common.PriceUpdateDto
	var totalDto common.TotalUpdateDto

	vegeName := "cucumber"

	// get RPC client by dialing at `rpc.DefaultRPCPath` endpoint
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	//Add new vege
	if err := client.Call("Shop.AddNewVegetable", common.Vegetable{Name: vegeName, Price: 15.3, AvailableTotal: 56}, &veg); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(veg.Price)
	}

	//Get available total
	if err := client.Call("Shop.GetAvailableTotalByName", vegeName, &availableAmount); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(availableAmount)
	}

	//Get Price
	if err := client.Call("Shop.GetPriceByName", vegeName, &price); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(price)
	}

	//Update available total
	if err := client.Call("Shop.UpdateAvailableTotalByName", common.TotalUpdateDto{
		Total: 56,
		Name:  vegeName,
	}, &totalDto); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(totalDto)
	}

	//Get Price
	if err := client.Call("Shop.UpdatePriceByName", common.PriceUpdateDto{
		Name:  vegeName,
		Price: 89.7,
	}, &priceDto); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(priceDto)
	}

	//Get All items
	if err := client.Call("Shop.GetAllNames", "", &allVeg); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(allVeg)
	}

}
