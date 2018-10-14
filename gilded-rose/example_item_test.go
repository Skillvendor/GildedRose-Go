// +build example

package gildedrose

import "fmt"

func ExampleGildedRose() {
	var items []*Item

	items = append(items, New("normal", 5, 5))
	items = append(items, New("Aged Brie", 5, 5))
	items = append(items, New("Sulfuras, Hand of Ragnaros", 5, 5))
	items = append(items, New("Backstage passes to a TAFKAL80ETC concert", 5, 5))
	// when you're ready to uncomment the line for the conjured item
	// you will also have to update the expected output below
	//
	// items = append(items, New("Conjured Mana Cake", 5, 5))

	for _, item := range items {
		fmt.Println(item)
	}

	UpdateQuality(items...)

	for _, item := range items {
		fmt.Println(item)
	}

	// Output:
	// normal: 5 days left, quality is 5
	// Aged Brie: 5 days left, quality is 5
	// Sulfuras, Hand of Ragnaros: 5 days left, quality is 5
	// Backstage passes to a TAFKAL80ETC concert: 5 days left, quality is 5
	// normal: 4 days left, quality is 4
	// Aged Brie: 4 days left, quality is 6
	// Sulfuras, Hand of Ragnaros: 5 days left, quality is 5
	// Backstage passes to a TAFKAL80ETC concert: 4 days left, quality is 8
}
