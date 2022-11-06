package autostore

type Product interface {
	Name() string
	Quantity() int
	Price() float64
	Sell(qty int)
	IsFinished() bool
}

// type product struct {
// 	id             productID
// 	name           string
// 	price          float64
// 	productType    ProductType
// 	category       string
// 	createdAt      *time.Time
// 	description    string
// }

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
