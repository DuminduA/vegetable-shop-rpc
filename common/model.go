package common

import (
	"errors"
)

type Vegetable struct {
	Name           string
	Price          float32
	AvailableTotal float32
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
	//file data is in memory to easy processing
	vegetables map[string]*Vegetable
}

func NewShop() *Shop {
	return &Shop{
		vegetables: make(map[string]*Vegetable),
	}
}

/*
	Get all the vegetables in the shop
*/
func (s *Shop) GetAllNames(payload string, result *[]string) error {
	//payload is not required here. just adding to support rpc lib
	//dat, err := ioutil.ReadFile("server/resources/database.txt")
	//check(err)
	//fmt.Print(string(dat))

	var allNames []string

	for idx, _ := range s.vegetables {
		allNames = append(allNames, idx)
	}

	result = &allNames
	return nil
}

func (s *Shop) GetPriceByName(name string, result *float32) error {
	current := s.vegetables[name]

	if current == nil {
		return errors.New("vegetable with this name is does not exist")
	}

	result = &current.Price
	return nil
}

func (s *Shop) GetAvailableTotalByName(name string, result *float32) error {
	current := s.vegetables[name]

	if current == nil {
		return errors.New("vegetable with this name is does not exist")
	}

	result = &current.AvailableTotal
	return nil
}

func (s *Shop) AddNewVegetable(vegetable Vegetable, result *Vegetable) error {
	current := s.vegetables[vegetable.Name]

	if current == nil {
		s.vegetables[vegetable.Name] = &vegetable
		result = s.vegetables[vegetable.Name]
	}

	return errors.New("vegetable with this name is does not exist")

}

func (s *Shop) UpdatePriceByName(payload PriceUpdateDto, result *Vegetable) error {
	current := s.vegetables[payload.name]

	if current == nil {
		return errors.New("vegetable with this name is does not exist")
	}
	current.Price = payload.price
	result = s.vegetables[payload.name]
	return nil

}

func (s *Shop) UpdateAvailableTotalByName(payload TotalUpdateDto, result *Vegetable) error {
	current := s.vegetables[payload.name]

	if current == nil {
		return errors.New("vegetable with this name is does not exist")
	}
	current.Price = payload.total
	result = s.vegetables[payload.name]
	return nil
}

/*
**************************************
		Helper structs
**************************************
*/

type PriceUpdateDto struct {
	name  string
	price float32
}

type TotalUpdateDto struct {
	name  string
	total float32
}
