package main

// Updater defines the behaviour of any object that can
// update its quality.
type Updater interface {
	UpdateQuality()
}

// Item holds information about a product.
type Item struct {
	name    string
	sellIn  int
	quality int
}

// BackstagePass is exactly what you'd think it is.
type BackstagePass struct {
	*Item
}

// UpdateQuality updates the quality of a BackstagePass.
func (item *BackstagePass) UpdateQuality() {
	item.quality++

	if item.sellIn < 11 {
		item.quality++
	}
	if item.sellIn < 6 {
		item.quality++
	}

	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		item.quality = 0
	}

	if item.quality > 50 {
		item.quality = 50
	}
}

// Aged item is any item whose value increases with age.
type Aged struct {
	*Item
}

// UpdateQuality updates the quality of an Aged item.
func (item *Aged) UpdateQuality() {
	item.quality++
	item.sellIn--

	if item.sellIn < 0 {
		item.quality++
	}

	if item.quality > 50 {
		item.quality = 50
	}
}

// Standard is an item whose quality is inversely correlated
// with its sell in value.
type Standard struct {
	*Item
}

// UpdateQuality updates the quality of a Standard item.
func (item *Standard) UpdateQuality() {
	item.quality--
	item.sellIn--

	if item.sellIn < 0 {
		item.quality--
	}

	if item.quality < 0 {
		item.quality = 0
	}
}

// Conjured item is exactly what you'd think it is.
type Conjured struct {
	*Item
}

// UpdateQuality updates the quality of a Conjured item.
func (item *Conjured) UpdateQuality() {
	item.quality -= 2

	item.sellIn--

	if item.sellIn < 0 {
		item.quality -= 2
	}

	if item.quality < 0 {
		item.quality = 0
	}
}

// Legendary is an item whose quality is not affected
// by its sell in value.
type Legendary struct {
	*Item
}

// UpdateQuality is a no-op on a Legendary item.
func (item *Legendary) UpdateQuality() {
	// noop, legendary items don't deteriorate
}
