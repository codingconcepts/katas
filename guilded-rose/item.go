package main

type Updater interface {
	UpdateQuality()
}

type Item struct {
	name    string
	sellIn  int
	quality int
}

type BackstagePass struct {
	Item
}

func (item *BackstagePass) UpdateQuality() {

}

type Aged struct {
	Item
}

func (item *Aged) UpdateQuality() {

}

type Standard struct {
	Item
}

func (item *Standard) UpdateQuality() {

}

type Conjured struct {
	Item
}

func (item *Conjured) UpdateQuality() {

}

type Legendary struct {
	Item
}

func (item *Legendary) UpdateQuality() {

}
