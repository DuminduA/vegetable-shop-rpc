package main

import (
	"fmt"
	"io/ioutil"
	"vegetable-shop-rpc/common"
)

func main() {

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
