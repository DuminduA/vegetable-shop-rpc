package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	common "vegetable-shop-rpc/server/model"
)

type VegetableController struct {
}

func newVegetableController() *VegetableController {
	return &VegetableController{}
}

func (vc VegetableController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/shop/getAll" {
		vc.GetAllNames(w, r)
	} else if r.URL.Path == "/shop/price" {
		switch r.Method {
		case http.MethodGet:
			vc.GetPrice(w, r)
		case http.MethodPut:
			vc.UpdatePrice(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else if r.URL.Path == "/shop/amount" {
		switch r.Method {
		case http.MethodGet:
			vc.GetAmount(w, r)
		case http.MethodPut:
			vc.UpdateAmount(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else if r.URL.Path == "/shop/add" {
		vc.AddNewVegetable(w, r)
	}

}

func (vc VegetableController) GetAllNames(w http.ResponseWriter, r *http.Request) {
	var allVeg []string

	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	if err := client.Call("Shop.GetAllNames", "", &allVeg); err != nil {
		fmt.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println(allVeg)
		encodeResponseAsJSON(allVeg, w)
	}
}

func (vc VegetableController) AddNewVegetable(w http.ResponseWriter, r *http.Request) {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	var veg common.Vegetable

	v, err := vc.parseRequestToVeg(r)

	if err != nil {
		fmt.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Add new vege
	if err := client.Call("Shop.AddNewVegetable", v, &veg); err != nil {
		fmt.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println(veg.Price)
		encodeResponseAsJSON(veg, w)
	}
}

func (vc VegetableController) UpdatePrice(w http.ResponseWriter, r *http.Request) {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	var priceDto common.PriceUpdateDto

	v, err := vc.parseRequestToPrice(r)

	if err != nil {
		fmt.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Get Price
	if err := client.Call("Shop.UpdatePriceByName", common.PriceUpdateDto{
		Name:  v.Name,
		Price: v.Price,
	}, &priceDto); err != nil {
		fmt.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println(priceDto)
		encodeResponseAsJSON(priceDto, w)
	}
}

func (vc VegetableController) UpdateAmount(w http.ResponseWriter, r *http.Request) {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	var totalDto common.TotalUpdateDto

	v, err := vc.parseRequestToAmount(r)

	if err != nil {
		fmt.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Update available total
	if err := client.Call("Shop.UpdateAvailableTotalByName", common.TotalUpdateDto{
		Total: v.Total,
		Name:  v.Name,
	}, &totalDto); err != nil {
		fmt.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println(totalDto)
		encodeResponseAsJSON(totalDto, w)
	}
}

func (vc VegetableController) GetPrice(w http.ResponseWriter, r *http.Request) {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	var price float32

	//Get Price
	if err := client.Call("Shop.GetPriceByName", r.FormValue("name"), &price); err != nil {
		fmt.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println(price)
		encodeResponseAsJSON(price, w)
	}
}

func (vc VegetableController) GetAmount(w http.ResponseWriter, r *http.Request) {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	var availableAmount float32

	//Get available total
	if err := client.Call("Shop.GetAvailableTotalByName", r.FormValue("name"), &availableAmount); err != nil {
		fmt.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println(availableAmount)
		encodeResponseAsJSON(availableAmount, w)
	}
}

/*
********************************************************************************************************
						Helper Methods
********************************************************************************************************
*/

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func (vc *VegetableController) parseRequestToVeg(r *http.Request) (common.Vegetable, error) {
	dec := json.NewDecoder(r.Body)
	var v common.Vegetable
	err := dec.Decode(&v)
	if err != nil {
		return common.Vegetable{}, err
	}
	return v, nil
}
func (vc *VegetableController) parseRequestToPrice(r *http.Request) (common.PriceUpdateDto, error) {
	dec := json.NewDecoder(r.Body)
	var v common.PriceUpdateDto
	err := dec.Decode(&v)
	if err != nil {
		return common.PriceUpdateDto{}, err
	}
	return v, nil
}

func (vc *VegetableController) parseRequestToAmount(r *http.Request) (common.TotalUpdateDto, error) {
	dec := json.NewDecoder(r.Body)
	var v common.TotalUpdateDto
	err := dec.Decode(&v)
	if err != nil {
		return common.TotalUpdateDto{}, err
	}
	return v, nil
}
