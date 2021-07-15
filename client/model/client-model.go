package common

type Vegetable struct {
	Name           string
	Price          float32
	AvailableTotal float32
}

/*
**************************************
		Helper structs
**************************************
*/

type PriceUpdateDto struct {
	Name  string
	Price float32
}

type TotalUpdateDto struct {
	Name  string
	Total float32
}
