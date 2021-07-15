package controller

import "net/http"

func RegisterControllers() {
	vc := newVegetableController()

	http.Handle("/shop", *vc)
	http.Handle("/shop/", *vc)
}
