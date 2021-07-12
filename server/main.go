package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/rpc"
	"vegetable-shop-rpc/common"
)

func main() {
	dat, err := ioutil.ReadFile("server/resources/database.txt")

	fmt.Println(err)
	fmt.Print(string(dat))

	// create a `*College` object
	shop := common.NewCollege()

	// register `mit` object with `rpc.DefaultServer`
	rpc.Register(shop)

	// register an HTTP handler for RPC communication on `http.DefaultServeMux` (default)
	// registers a handler on the `rpc.DefaultRPCPath` endpoint to respond to RPC messages
	// registers a handler on the `rpc.DefaultDebugPath` endpoint for debugging
	rpc.HandleHTTP()

	// sample test endpoint
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPC SERVER LIVE!")
	})

	// listen and serve default HTTP server
	http.ListenAndServe(":9000", nil)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetAllNames() []common.Vegetable {
	dat, err := ioutil.ReadFile("server/resources/database.txt")
	check(err)
	fmt.Print(string(dat))

	return nil
}

func GetPriceByName(Name string) (Price float32, err error) {
	return 0, err
}

func GetAvailableTotalByName(Name string) (Price float32, err error) {
	return 0, err
}

func AddNewVegetable(vege common.Vegetable) (vegeSaved common.Vegetable, err error) {
	return common.Vegetable{}, err
}

func UpdatePriceByName(Name string, Price float32) (vegeSaved common.Vegetable, err error) {
	return common.Vegetable{}, err
}

func UpdateAvailableTotalByName(Name string, Price float32) (vegeSaved common.Vegetable, err error) {
	return common.Vegetable{}, err
}

