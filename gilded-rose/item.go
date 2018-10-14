package gildedrose

import "fmt"

// Item describes an item sold by the Gilded Rose Inn
type Item struct {
	name    string
	days    int
	quality int
}

// String outputs a string representation of an Item
func (i *Item) String() string {
	return fmt.Sprintf("%s: %d days left, quality is %d", i.name, i.days, i.quality)
}

// New creates a new Item
func New(name string, days, quality int) *Item {
	return &Item{
		name:    name,
		days:    days,
		quality: quality,
	}
}

// UpdateQuality ages the item by a day, and updates the quality of the item
func UpdateQuality(items ...*Item) {
	for _, item := range items {
		switch item.name {
		case "normal":
			normalTick(item)
		case "Aged Brie":
			agedBrieTick(item)
		case "Sulfuras, Hand of Ragnaros":
			sulfurasTick(item)
		case "Backstage passes to a TAFKAL80ETC concert":
			backstageTick(item)
		default:
			return
		}
	}
}

func normalTick(item *Item) {
	item.days--
	if item.quality == 0 {
		return
	}

	item.quality--
	if item.days <= 0 {
		item.quality--
	}
}

func agedBrieTick(item *Item) {
	item.days--
	if item.quality >= 50 {
		return
	}

	if item.days <= 0 {
		item.quality++
	}
	if item.quality < 50 {
		item.quality++
	}
}

func sulfurasTick(item *Item) {

}

func backstageTick(item *Item) {
	item.days--
	if item.quality >= 50 {
		return
	}
	if item.days < 0 {
		item.quality = 0
		return
	}

	item.quality++
	if item.days < 10 && item.quality < 50 {
		item.quality++
	}
	if item.days < 5 && item.quality < 50 {
		item.quality++
	}
}
