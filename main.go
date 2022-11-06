package autostore

import (
	"fmt"
)

func main() {
	fmt.Println("YOU ARE HERE, THE BEST AUTO SHOP")
	Newline(2)
	store := NewStore()

	fmt.Println("Creating cars .......")
	Gwagon := AddCar("Mercedes", "GWagon", "Red", 50, 340000000)
	x5 := AddCar("BWMW", "x5", "black", 5, 78000000)
	highlander := AddCar("Toyota", "highlander", "blue", 10, 9000000)
	fmt.Println("Items are being added to the store..........")
	store.AddItem(&highlander)
	store.AddItem(&Gwagon)
	store.AddItem(&x5)
	store.NumberOfProduct()
	fmt.Println("Items available in store")
	store.DisplayAllItem()
	fmt.Println("Selling...")
	Newline(1)
	store.Sell(2, 6)
	store.Sell(1, 4)
	store.Sell(3, 15)
	store.Sell(4, 2)
	fmt.Println("Items available after selling")
	store.DisplayAllItem()
	store.ItemsLog()
}
