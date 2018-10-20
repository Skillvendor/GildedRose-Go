package gildedrose

import "fmt"

// Ticker is the interface that wraps the Update method.
type Ticker interface {
	// Tick updates the Quality and Days remaining.
	Tick()
}

// Item describes an item sold by the Gilded Rose Inn
type Item struct {
	Name    string
	Days    int
	Quality int
}

// Tick updates the Quality and Days remaining.
func (item Item) Tick() {
	// do nothing
}

// String outputs a string representation of an Item
func (i *Item) String() string {
	return fmt.Sprintf("%s: %d days left, quality is %d", i.Name, i.Days, i.Quality)
}

// Normal is a normal item.
type Normal struct {
	*Item
}

// Tick updates the Quality and Days remaining.
func (item Normal) Tick() {
	item.Days--
	if item.Quality == 0 {
		return
	}

	item.Quality--
	if item.Days <= 0 {
		item.Quality--
	}
}

// Brie is 'Aged Brie'.
type Brie struct {
	*Item
}

// Tick updates the Quality and Days remaining.
func (item Brie) Tick() {
	item.Days--
	if item.Quality >= 50 {
		return
	}

	if item.Days <= 0 {
		item.Quality++
	}
	if item.Quality < 50 {
		item.Quality++
	}
}

type Backstage struct {
	*Item
}

func (item Backstage) Tick() {
	item.Days--
	if item.Quality >= 50 {
		return
	}
	if item.Days < 0 {
		item.Quality = 0
		return
	}

	item.Quality++
	if item.Days < 10 && item.Quality < 50 {
		item.Quality++
	}
	if item.Days < 5 && item.Quality < 50 {
		item.Quality++
	}
}

type Conjured struct {
	*Item
}

func (item Conjured) Tick() {
	item.Days--
	if item.Quality == 0 {
		return
	}

	item.Quality -= 2
	if item.Days < 0 {
		item.Quality -= 2
	}
}

// UpdateQuality ages the item by a day, and updates the quality of the item
func UpdateQuality(items ...*Item) {
	for _, item := range items {
		var t Ticker
		switch item.Name {
		case "normal":
			t = Normal{item}
		case "Aged Brie":
			t = Brie{item}
		case "Backstage passes to a TAFKAL80ETC concert":
			t = Backstage{item}
		case "Conjured Mana Cake":
			t = Conjured{item}
		default:
			t = item
		}
		t.Tick()
	}
}
