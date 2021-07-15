package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	common "vegetable-shop-rpc/server/model"
)

func main() {
	//read the file
	//dat, err := ioutil.ReadFile("server/resources/database.txt")

	//initialized the server as in the reference https://medium.com/rungo/building-rpc-remote-procedure-call-network-in-go-5bfebe90f7e9
	shop := common.NewShop()

	rpc.Register(shop)

	// register an HTTP handler for RPC communication on `http.DefaultServeMux` (default)
	// registers a handler on the `rpc.DefaultRPCPath` endpoint to respond to RPC messages
	// registers a handler on the `rpc.DefaultDebugPath` endpoint for debugging
	rpc.HandleHTTP()

	// sample test endpoint
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPC SERVER LIVE!")
	})

	fmt.Println("Server started on port 9000")

	// listen and serve default HTTP server
	http.ListenAndServe(":9000", nil)

}
