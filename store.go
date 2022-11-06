package main

import (
	"fmt"

	guuid "github.com/google/uuid"
)

func Newline(num int) {
	for num > 0 {
		fmt.Println()
		num--
	}
}

func GenUUID() string {
	id := guuid.New()
	Id := fmt.Sprintf("%v", id)
	return Id[:6]
}

func AddCar(brandName, model, color string, quantity int, price float64) car {
	id := GenUUID()
	return car{
		id:        id,
		brandName: brandName,
		model:     model,
		color:     color,
		quantity:  quantity,
		price:     price,
	}
}

func (c *car) Name() string {
	return fmt.Sprintf("%s-%s [%s]", c.brandName, c.model, c.color)
}

func (c *car) Quantity() int {
	return c.quantity
}

func (c *car) Price() float64 {
	return c.price
}

func (c *car) Sell(num int) {
	c.quantity = c.quantity - num
}

func (c *car) IsFinished() bool {
	return c.quantity > 1
}

func Display(p Product) {
	name := p.Name()
	quantity := p.Quantity()
	price := p.Price()

	fmt.Printf("Name - %s \nQuantity - %v \nPrice - %.2f", name, quantity,
		price)
	Newline(2)
}

var logBook []Order // store data

func NewStore() *Store {
	return &Store{
		myStore: make(map[int]Product),
		Orders:  []Order{},
	}
}

func (s Store) NumberOfProduct() {
	num := 0
	for i := 0; i < len(s.myStore); i++ {
		num++
	}

	fmt.Printf("You have %v types of products in this autostore", num)
	Newline(2)
}

func (s Store) AddItem(item Product) {
	num := len(s.myStore)
	num++
	s.myStore[num] = item
}

func (s Store) DisplayAllItem() {
	fmt.Println("Below is a list of all the items in the autostore")
	Newline(1)
	fmt.Println("= ID ====== Name ======== Price === Quantity")
	var num_of_cars_left int
	var total_sum_of_cars_left float64
	for i, v := range s.myStore {
		total_sum_of_cars_left += float64(v.Quantity()) * v.Price()
		num_of_cars_left += v.Quantity()
		fmt.Printf("%v - %s - %.2f - %v", i, v.Name(), v.Price(), v.Quantity())
		Newline(1)
	}

	Newline(2)
	fmt.Printf("Total number of cars left  =>  %d\n", num_of_cars_left)
	fmt.Printf("Total sum total of cars left  =>  %.2f", total_sum_of_cars_left)
	Newline(3)
}

func (s Store) Sell(itemId, qty int) {
	item := s.myStore[itemId]
	name := item.Name()
	if !item.IsFinished() {
		fmt.Printf("No item %s in stock", name)
	} else if item.Quantity() < qty {
		fmt.Println("Quantity left in stock is less than you requested \nquantity left ", item.Quantity())
	} else {
		item.Sell(qty)
		newOrder := Order{
			itemId:   itemId,
			quantity: qty,
		}
		s.Orders = append(s.Orders, newOrder)
	}

	logBook = append(logBook, s.Orders...)
}

func (s Store) ItemsLog() {
	Newline(1)
	fmt.Println("HISTORY OF SALES")
	fmt.Println("Num ====== Name ===== Price ==== Qty = Total price")
	var total_sales_made float64
	var total_number_of_cars_sold int
	for i, v := range logBook {
		num := i + 1
		item := s.myStore[v.itemId]
		price := item.Price()
		qty := v.quantity
		totalPrice := float64(qty) * price

		fmt.Printf("%d - %v - %.2f, %v, %.2f", num, item.Name(), price, qty, totalPrice)
		total_sales_made += totalPrice
		total_number_of_cars_sold += qty
		Newline(1)
	}

	Newline(1)
	fmt.Printf("Total number of cars sold  =>  %d\n", total_number_of_cars_sold)
	fmt.Printf("Total sum of cars sold  =>  %.2f", total_sales_made)
}
