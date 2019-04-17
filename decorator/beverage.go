package main

type Size int

const (
	_ Size = iota
	TALL
	GRANDE
	VENTI
)

// Beverage ---
type Beverage interface {
	GetDescription() string
	SetSize(s Size)
	GetSize() Size
	Cost() float64
}

type beverage struct {
	size        Size
	description string
}

func (b *beverage) GetDescription() string {
	return b.description
}

func (b *beverage) SetSize(s Size) {
	b.size = s
}

func (b *beverage) GetSize() Size {
	return b.size
}

// DarkRoast ---
type DarkRoast struct {
	beverage
}

func NewDarkRoast() *DarkRoast {
	dr := &DarkRoast{}
	dr.description = "Dark Roast Coffee"

	return dr
}

func (dr *DarkRoast) Cost() float64 {
	return .99
}

// Decaf ---
type Decaf struct {
	beverage
}

func NewDecaf() *Decaf {
	d := &Decaf{}
	d.description = "Decaf Coffee"

	return d
}

func (d *Decaf) Cost() float64 {
	return 1.05
}

// Espresso ---
type Espresso struct {
	beverage
}

func NewEspresso() *Espresso {
	e := &Espresso{}
	e.description = "Espresso"

	return e
}

func (e *Espresso) Cost() float64 {
	return 1.99
}

// HouseBland ---
type HouseBland struct {
	beverage
}

func NewHouseBland() *HouseBland {
	hb := &HouseBland{}
	hb.description = "House Bland Coffee"

	return hb
}

func (hb *HouseBland) Cost() float64 {
	return .89
}
