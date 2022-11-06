package main

type Product interface {
	Name() string
	Quantity() int
	Price() float64
	Sell(qty int)
	IsFinished() bool
}

type Order struct {
	itemId   int
	quantity int
}

type Store struct {
	myStore map[int]Product
	Orders  []Order
}

type car struct {
	id        string
	brandName string
	model     string
	color     string
	quantity  int
	price     float64
}
