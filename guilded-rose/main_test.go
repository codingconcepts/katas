package main

import (
	"fmt"
	"testing"
)

func Test_DeterityVest(t *testing.T) {
	testCases := []struct {
		act Item
		exp Item
	}{
		{
			act: Item{name: "+5 Dexterity Vest", sellIn: 10, quality: 20},
			exp: Item{name: "+5 Dexterity Vest", sellIn: 9, quality: 19},
		},
		{
			act: Item{name: "+5 Dexterity Vest", sellIn: 9, quality: 20},
			exp: Item{name: "+5 Dexterity Vest", sellIn: 8, quality: 19},
		},
		{
			act: Item{name: "+5 Dexterity Vest", sellIn: 1, quality: 20},
			exp: Item{name: "+5 Dexterity Vest", sellIn: 0, quality: 19},
		},
		{
			act: Item{name: "+5 Dexterity Vest", sellIn: 0, quality: 20},
			exp: Item{name: "+5 Dexterity Vest", sellIn: -1, quality: 18},
		},
		{
			act: Item{name: "+5 Dexterity Vest", sellIn: 0, quality: 1},
			exp: Item{name: "+5 Dexterity Vest", sellIn: -1, quality: 0},
		},
		{
			act: Item{name: "+5 Dexterity Vest", sellIn: 0, quality: 0},
			exp: Item{name: "+5 Dexterity Vest", sellIn: -1, quality: 0},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_DeterityVest_%v", testCase.act), func(t *testing.T) {
			items := []Item{testCase.act}
			UpdateQuality(items)

			first := items[0]

			if first.quality != testCase.exp.quality {
				t.Fatalf("quality exp: '%d' but got '%d'", testCase.exp.quality, first.quality)
			}

			if first.sellIn != testCase.exp.sellIn {
				t.Fatalf("sellIn exp: '%d' but got '%d'", testCase.exp.sellIn, first.sellIn)
			}
		})
	}
}

func Test_AgedBrie(t *testing.T) {
	testCases := []struct {
		act Item
		exp Item
	}{
		{
			act: Item{name: "Aged Brie", sellIn: 10, quality: 20},
			exp: Item{name: "Aged Brie", sellIn: 9, quality: 21},
		},
		{
			act: Item{name: "Aged Brie", sellIn: 10, quality: 50},
			exp: Item{name: "Aged Brie", sellIn: 9, quality: 50},
		},
		{
			act: Item{name: "Aged Brie", sellIn: 10, quality: 49},
			exp: Item{name: "Aged Brie", sellIn: 9, quality: 50},
		},
		{
			act: Item{name: "Aged Brie", sellIn: 9, quality: 20},
			exp: Item{name: "Aged Brie", sellIn: 8, quality: 21},
		},
		{
			act: Item{name: "Aged Brie", sellIn: 1, quality: 20},
			exp: Item{name: "Aged Brie", sellIn: 0, quality: 21},
		},
		{
			act: Item{name: "Aged Brie", sellIn: 0, quality: 20},
			exp: Item{name: "Aged Brie", sellIn: -1, quality: 22},
		},
		{
			act: Item{name: "Aged Brie", sellIn: 0, quality: 50},
			exp: Item{name: "Aged Brie", sellIn: -1, quality: 50},
		},
		{
			act: Item{name: "Aged Brie", sellIn: 0, quality: 49},
			exp: Item{name: "Aged Brie", sellIn: -1, quality: 50},
		},
		{
			act: Item{name: "Aged Brie", sellIn: 0, quality: 48},
			exp: Item{name: "Aged Brie", sellIn: -1, quality: 50},
		},
		{
			act: Item{name: "Aged Brie", sellIn: 0, quality: 1},
			exp: Item{name: "Aged Brie", sellIn: -1, quality: 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_AgedBrie_%v", testCase.act), func(t *testing.T) {
			items := []Item{testCase.act}
			UpdateQuality(items)

			first := items[0]

			if first.quality != testCase.exp.quality {
				t.Fatalf("quality exp: '%d' but got '%d'", testCase.exp.quality, first.quality)
			}

			if first.sellIn != testCase.exp.sellIn {
				t.Fatalf("sellIn exp: '%d' but got '%d'", testCase.exp.sellIn, first.sellIn)
			}
		})
	}
}

func Test_Elixir(t *testing.T) {
	testCases := []struct {
		act Item
		exp Item
	}{
		{
			act: Item{name: "Elixir of the Mongoose", sellIn: 10, quality: 20},
			exp: Item{name: "Elixir of the Mongoose", sellIn: 9, quality: 19},
		},
		{
			act: Item{name: "Elixir of the Mongoose", sellIn: 9, quality: 20},
			exp: Item{name: "Elixir of the Mongoose", sellIn: 8, quality: 19},
		},
		{
			act: Item{name: "Elixir of the Mongoose", sellIn: 1, quality: 20},
			exp: Item{name: "Elixir of the Mongoose", sellIn: 0, quality: 19},
		},
		{
			act: Item{name: "Elixir of the Mongoose", sellIn: 0, quality: 20},
			exp: Item{name: "Elixir of the Mongoose", sellIn: -1, quality: 18},
		},
		{
			act: Item{name: "Elixir of the Mongoose", sellIn: 0, quality: 1},
			exp: Item{name: "Elixir of the Mongoose", sellIn: -1, quality: 0},
		},
		{
			act: Item{name: "Elixir of the Mongoose", sellIn: 0, quality: 0},
			exp: Item{name: "Elixir of the Mongoose", sellIn: -1, quality: 0},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_Elixir_%v", testCase.act), func(t *testing.T) {
			items := []Item{testCase.act}
			UpdateQuality(items)

			first := items[0]

			if first.quality != testCase.exp.quality {
				t.Fatalf("quality exp: '%d' but got '%d'", testCase.exp.quality, first.quality)
			}

			if first.sellIn != testCase.exp.sellIn {
				t.Fatalf("sellIn exp: '%d' but got '%d'", testCase.exp.sellIn, first.sellIn)
			}
		})
	}
}

func Test_Sulfuras(t *testing.T) {
	testCases := []struct {
		act Item
		exp Item
	}{
		{
			act: Item{name: "Sulfuras, Hand of Ragnaros", sellIn: 10, quality: 20},
			exp: Item{name: "Sulfuras, Hand of Ragnaros", sellIn: 10, quality: 20},
		},
		{
			act: Item{name: "Sulfuras, Hand of Ragnaros", sellIn: 9, quality: 20},
			exp: Item{name: "Sulfuras, Hand of Ragnaros", sellIn: 9, quality: 20},
		},
		{
			act: Item{name: "Sulfuras, Hand of Ragnaros", sellIn: 1, quality: 20},
			exp: Item{name: "Sulfuras, Hand of Ragnaros", sellIn: 1, quality: 20},
		},
		{
			act: Item{name: "Sulfuras, Hand of Ragnaros", sellIn: 0, quality: 20},
			exp: Item{name: "Sulfuras, Hand of Ragnaros", sellIn: 0, quality: 20},
		},
		{
			act: Item{name: "Sulfuras, Hand of Ragnaros", sellIn: 0, quality: 1},
			exp: Item{name: "Sulfuras, Hand of Ragnaros", sellIn: 0, quality: 1},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_Sulfuras_%v", testCase.act), func(t *testing.T) {
			items := []Item{testCase.act}
			UpdateQuality(items)

			first := items[0]

			if first.quality != testCase.exp.quality {
				t.Fatalf("quality exp: '%d' but got '%d'", testCase.exp.quality, first.quality)
			}

			if first.sellIn != testCase.exp.sellIn {
				t.Fatalf("sellIn exp: '%d' but got '%d'", testCase.exp.sellIn, first.sellIn)
			}
		})
	}
}

func Test_BackstagePass(t *testing.T) {
	testCases := []struct {
		act Item
		exp Item
	}{
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 11, quality: 20},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 21},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 11, quality: 50},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 11, quality: 49},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 20},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 9, quality: 22},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 50},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 9, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 49},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 9, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 48},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 9, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 6, quality: 20},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 22},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 6, quality: 50},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 6, quality: 49},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 6, quality: 48},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 20},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 23},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 50},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 49},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 48},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 47},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 50},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 9, quality: 20},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 8, quality: 22},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 1, quality: 20},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 0, quality: 23},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 0, quality: 20},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: -1, quality: 0},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 0, quality: 1},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: -1, quality: 0},
		},
		{
			act: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 0, quality: 0},
			exp: Item{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: -1, quality: 0},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_BackstagePass_%v", testCase.act), func(t *testing.T) {
			items := []Item{testCase.act}
			UpdateQuality(items)

			first := items[0]

			if first.quality != testCase.exp.quality {
				t.Fatalf("quality exp: '%d' but got '%d'", testCase.exp.quality, first.quality)
			}

			if first.sellIn != testCase.exp.sellIn {
				t.Fatalf("sellIn exp: '%d' but got '%d'", testCase.exp.sellIn, first.sellIn)
			}
		})
	}
}

// TODO: update to reflect change in system post refactor
func Test_Conjured(t *testing.T) {
	testCases := []struct {
		act Item
		exp Item
	}{
		{
			act: Item{name: "Conjured Mana Cake", sellIn: 10, quality: 20},
			exp: Item{name: "Conjured Mana Cake", sellIn: 9, quality: 19},
		},
		{
			act: Item{name: "Conjured Mana Cake", sellIn: 9, quality: 20},
			exp: Item{name: "Conjured Mana Cake", sellIn: 8, quality: 19},
		},
		{
			act: Item{name: "Conjured Mana Cake", sellIn: 1, quality: 20},
			exp: Item{name: "Conjured Mana Cake", sellIn: 0, quality: 19},
		},
		{
			act: Item{name: "Conjured Mana Cake", sellIn: 0, quality: 20},
			exp: Item{name: "Conjured Mana Cake", sellIn: -1, quality: 18},
		},
		{
			act: Item{name: "Conjured Mana Cake", sellIn: 0, quality: 1},
			exp: Item{name: "Conjured Mana Cake", sellIn: -1, quality: 0},
		},
		{
			act: Item{name: "Conjured Mana Cake", sellIn: 0, quality: 0},
			exp: Item{name: "Conjured Mana Cake", sellIn: -1, quality: 0},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_Conjured_%v", testCase.act), func(t *testing.T) {
			items := []Item{testCase.act}
			UpdateQuality(items)

			first := items[0]

			if first.quality != testCase.exp.quality {
				t.Fatalf("quality exp: '%d' but got '%d'", testCase.exp.quality, first.quality)
			}

			if first.sellIn != testCase.exp.sellIn {
				t.Fatalf("sellIn exp: '%d' but got '%d'", testCase.exp.sellIn, first.sellIn)
			}
		})
	}
}
