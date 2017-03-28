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
	if item.quality < 50 {
		item.quality++
		if item.sellIn < 11 {
			if item.quality < 50 {
				item.quality++
			}
		}
		if item.sellIn < 6 {
			if item.quality < 50 {
				item.quality++
			}
		}
	}

	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		item.quality = 0
	}
}

// Aged item is any item whose value increases with age.
type Aged struct {
	*Item
}

// UpdateQuality updates the quality of an Aged item.
func (item *Aged) UpdateQuality() {
	if item.quality < 50 {
		item.quality++
	}

	item.sellIn--

	if item.sellIn >= 0 {
		return
	}

	if item.quality < 50 {
		item.quality++
	}
}

// Standard is an item whose quality is inversely correlated
// with its sell in value.
type Standard struct {
	*Item
}

// UpdateQuality updates the quality of a Standard item.
func (item *Standard) UpdateQuality() {
	if item.quality > 0 {
		item.quality--
	}

	item.sellIn--

	if item.sellIn < 0 {
		if item.quality > 0 {
			item.quality--
		}
	}
}

// Conjured item is exactly what you'd think it is.
type Conjured struct {
	*Item
}

// UpdateQuality updates the quality of a Conjured item.
func (item *Conjured) UpdateQuality() {

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
