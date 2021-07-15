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

	*result = allNames
	return nil
}

func (s *Shop) GetPriceByName(name string, result *float32) error {
	current := s.vegetables[name]

	if current == nil {
		return errors.New("vegetable with this name is does not exist")
	}

	*result = current.Price
	return nil
}

func (s *Shop) GetAvailableTotalByName(name string, result *float32) error {
	current := s.vegetables[name]

	if current == nil {
		return errors.New("vegetable with this name is does not exist")
	}

	*result = current.AvailableTotal
	return nil
}

func (s *Shop) AddNewVegetable(vegetable Vegetable, result *Vegetable) error {
	current := s.vegetables[vegetable.Name]

	if current == nil {
		s.vegetables[vegetable.Name] = &vegetable
		*result = vegetable
		return nil
	}

	return errors.New("vegetable with this name is exists")

}

func (s *Shop) UpdatePriceByName(payload PriceUpdateDto, result *PriceUpdateDto) error {
	current := s.vegetables[payload.Name]

	if current == nil {
		return errors.New("vegetable with this name is does not exist")
	}
	current.Price = payload.Price
	s.vegetables[payload.Name] = current
	*result = PriceUpdateDto{Name: payload.Name, Price: payload.Price}
	return nil

}

func (s *Shop) UpdateAvailableTotalByName(payload TotalUpdateDto, result *TotalUpdateDto) error {
	current := s.vegetables[payload.Name]

	if current == nil {
		return errors.New("vegetable with this name is does not exist")
	}
	current.Price = payload.Total
	s.vegetables[payload.Name] = current
	*result = TotalUpdateDto{Name: payload.Name, Total: payload.Total}
	return nil
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
