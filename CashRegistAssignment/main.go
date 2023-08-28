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

func (item Item) Describe() string {
	if item.DiscountPrice == 0 {
		return fmt.Sprintf("%s - %.1f TL", item.Name, item.Price)
	}
	return fmt.Sprintf("%s - %.1f TL (İndirimle %.1f TL)", item.Name, item.Price, calcutePrice(item))
}
func calcutePrice(item Item) float64 {
	return item.Price - item.DiscountPrice
}

func totalPrice(items []Item) float64 {
	var total float64
	for _, item := range items {
		total += calcutePrice(item)
	}
	return total
}
func printTotal(items []Item) {
	fmt.Println("Toplam Tutar:", totalPrice(items), "TL")
}
func main() {
	elma := Item{
		Name:          "elma",
		Price:         10,
		DiscountPrice: 1.5,
	}
	armut := Item{
		Name:  "armut",
		Price: 19.5,
	}
	fmt.Println(elma.Describe())
	fmt.Println(armut.Describe())
	printTotal([]Item{elma, armut})
}
