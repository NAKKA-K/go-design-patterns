package main

// CondimentDecorator ---
type CondimentDecorator interface {
	Beverage
}

type condimentDecorator struct {
	beverage
	Beverage Beverage
}

func (cd *condimentDecorator) GetSize() Size {
	return cd.GetSize()
}

// Milk ---
type Milk struct {
	condimentDecorator
}

func NewMilk(beverage Beverage) *Milk {
	m := &Milk{}
	m.Beverage = beverage

	return m
}

func (m *Milk) GetDescription() string {
	return m.Beverage.GetDescription() + ", Milk"
}

func (m *Milk) Cost() float64 {
	return m.Beverage.Cost() + .10
}

// Mocha ---
type Mocha struct {
	condimentDecorator
}

func NewMocha(beverage Beverage) *Mocha {
	m := &Mocha{}
	m.Beverage = beverage

	return m
}

func (m *Mocha) GetDescription() string {
	return m.Beverage.GetDescription() + ", Mocha"
}

func (m *Mocha) Cost() float64 {
	return m.Beverage.Cost() + .20
}

// Soy ---
type Soy struct {
	condimentDecorator
}

func NewSoy(beverage Beverage) *Soy {
	m := &Soy{}
	m.Beverage = beverage

	return m
}

func (m *Soy) GetDescription() string {
	return m.Beverage.GetDescription() + ", Soy"
}

func (m *Soy) Cost() float64 {
	cost := m.Beverage.Cost()
	switch m.Beverage.GetSize() {
	case TALL:
		cost += .10
	case GRANDE:
		cost += .15
	case VENTI:
		cost += .20
	}
	return cost
}

// Whip ---
type Whip struct {
	condimentDecorator
}

func NewWhip(beverage Beverage) *Whip {
	m := &Whip{}
	m.Beverage = beverage

	return m
}

func (m *Whip) GetDescription() string {
	return m.Beverage.GetDescription() + ", Whip"
}

func (m *Whip) Cost() float64 {
	return m.Beverage.Cost() + .10
}
