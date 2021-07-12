package common

import (
	"fmt"
	"io/ioutil"
)

type Vegetable struct {
	Name string,
	Price float,
	AvailableTotal float
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
************************************
Vegetable shop
************************************
 */

type Shop struct {
	vegetables map[int]*Vegetable
}

func NewShop() *Shop {
	return &Shop{
		vegetables: make(map[int]*Vegetable),
	}
}
/*
	Get all the vegetables in the shop
 */
func GetAllNames() []Vegetable {
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

func AddNewVegetable(vege Vegetable) (vegeSaved Vegetable, err error) {
	return Vegetable{}, err
}

func UpdatePriceByName(Name string, Price float32) (vegeSaved Vegetable, err error) {
	return Vegetable{}, err
}

func UpdateAvailableTotalByName(Name string, Price float32) (vegeSaved Vegetable, err error) {
	return Vegetable{}, err
}
