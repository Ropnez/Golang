package main

import "fmt"

type Describable interface {
	Describe() string
}

type Item struct {
	Name          string
	Price         float64
	DiscountPrice float64
}

func calcutePrice(item Item) float64 {
	return item.Price - item.DiscountPrice
}

func totalPrice(items []Item) float64 {
	if len(items) == 0 {
		return 0
	}
	var total float64 = 0
	for _, item := range items {
		total += calcutePrice(item)
	}
	return total
}

func main() {
	a := Item{
		Name:          "elma",
		Price:         10,
		DiscountPrice: 1.5,
	}
	b := Item{
		Name:  "armut",
		Price: 19.5,
	}
	calcutePrice(a)
	fmt.Println(a.Name, "-", a.Price, "TL", "indirimli fiyatı:", calcutePrice(a), "TL")
	totalPrice([]Item{a})
	fmt.Println(b.Name, b.Price, "TL")
	fmt.Println("İndirimli toplam fiyat:", totalPrice([]Item{(a), b}), "TL")
}
