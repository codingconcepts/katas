package main

import (
	"fmt"
	"log"
	"os"
)

var items = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	Item{"Conjured Mana Cake", 3, 6},
}

var (
	mappings ItemMappings
)

func main() {
	// load mappings from file (or database etc.)
	file, err := os.Open("items.json")
	if err != nil {
		log.Fatal(err)
	}

	if mappings, err = LoadMappings(file); err != nil {
		log.Fatal(err)
	}

	UpdateQuality(items)

	for _, item := range items {
		fmt.Println(item)
	}
}

// UpdateQuality loops through a given collection of items
// and modifies their quality/sell in according to their
// type.
func UpdateQuality(items []Item) {
	for i := 0; i < len(items); i++ {
		updater := mappings.CategoriseItem(&items[i])
		updater.UpdateQuality()
	}
}
