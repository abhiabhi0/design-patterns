package factorymethod

import "fmt"

// Button is the abstract product
type Button interface {
	Render()
	OnClick()
	GetBorder() float64
	SetBorder(border float64)
}

// RoundButton is a concrete product
type RoundButton struct {
	border float64
	radius float64
}

func (b *RoundButton) Render() {
	fmt.Println("Rendered RoundButton!")
}

func (b *RoundButton) OnClick() {
	fmt.Println("Round Button was clicked!")
}

func (b *RoundButton) GetBorder() float64 {
	return b.border
}

func (b *RoundButton) SetBorder(border float64) {
	b.border = border
}

func (b *RoundButton) GetRadius() float64 {
	return b.radius
}

func (b *RoundButton) SetRadius(radius float64) {
	b.radius = radius
}

// SquareButton is a concrete product
type SquareButton struct {
	border float64
	length float64
}

func (b *SquareButton) Render() {
	fmt.Println("Rendered SquareButton!")
}

func (b *SquareButton) OnClick() {
	fmt.Println("Square Button was clicked!")
}

func (b *SquareButton) GetBorder() float64 {
	return b.border
}

func (b *SquareButton) SetBorder(border float64) {
	b.border = border
}

func (b *SquareButton) GetLength() float64 {
	return b.length
}

func (b *SquareButton) SetLength(length float64) {
	b.length = length
}

// ButtonFactory is the abstract factory
type ButtonFactory interface {
	CreateButton(border float64, radius float64, length float64) Button
}

// RoundButtonFactory is a concrete factory
type RoundButtonFactory struct{}

func (f *RoundButtonFactory) CreateButton(border float64, radius float64, length float64) Button {
	return &RoundButton{
		border: border,
		radius: radius,
	}
}

// SquareButtonFactory is a concrete factory
type SquareButtonFactory struct{}

func (f *SquareButtonFactory) CreateButton(border float64, radius float64, length float64) Button {
	return &SquareButton{
		border: border,
		length: length,
	}
}

func ClientCode(factory ButtonFactory) {
	button := factory.CreateButton(1.0, 5.0, 10.0)
	button.Render()
	button.OnClick()
}
