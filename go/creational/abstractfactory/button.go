package abstractfactory

import "fmt"

// Button is the abstract product
type Button interface {
	Render()
	OnClick()
	GetBorder() float64
}

// DarkButton is a concrete product
type DarkButton struct {
	border float64
	radius float64
}

func (b *DarkButton) Render() {
	fmt.Println("Rendered DarkButton!")
}

func (b *DarkButton) OnClick() {
	fmt.Println("Dark Btn was clicked!")
}

func (b *DarkButton) GetBorder() float64 {
	return b.border
}

func (b *DarkButton) GetRadius() float64 {
	return b.radius
}

// Radio is the abstract product
type Radio interface {
	OnSelect()
	Render()
}

// DarkRadio is a concrete product
type DarkRadio struct{}

func (r *DarkRadio) OnSelect() {
	fmt.Println("DarkRadio selected!")
}

func (r *DarkRadio) Render() {
	fmt.Println("DarkRadio rendered!")
}

// LightButton is a concrete product
type LightButton struct {
	border float64
	length float64
}

func (b *LightButton) Render() {
	fmt.Println("Rendered LightButton!")
}

func (b *LightButton) OnClick() {
	fmt.Println("Light Btn was clicked!")
}

func (b *LightButton) GetBorder() float64 {
	return b.border
}

func (b *LightButton) GetLength() float64 {
	return b.length
}

// LightRadio is a concrete product
type LightRadio struct{}

func (r *LightRadio) OnSelect() {
	fmt.Println("LightRadio selected!")
}

func (r *LightRadio) Render() {
	fmt.Println("LightRadio rendered!")
}

// ThemeFactory is the abstract factory
type ThemeFactory interface {
	CreateButton(border, length, radius float64) Button
	CreateRadio() Radio
}

// DarkThemeFactory is a concrete factory
type DarkThemeFactory struct{}

func (f *DarkThemeFactory) CreateButton(border, length, radius float64) Button {
	return &DarkButton{border: border, radius: radius}
}

func (f *DarkThemeFactory) CreateRadio() Radio {
	return &DarkRadio{}
}

// LightThemeFactory is a concrete factory
type LightThemeFactory struct{}

func (f *LightThemeFactory) CreateButton(border, length, radius float64) Button {
	return &LightButton{border: border, length: length}
}

func (f *LightThemeFactory) CreateRadio() Radio {
	return &LightRadio{}
}

// Client code
func ClientCode(factory ThemeFactory) {
	button := factory.CreateButton(1.0, 10.0, 5.0)
	button.Render()
	button.OnClick()

	radio := factory.CreateRadio()
	radio.Render()
	radio.OnSelect()
}
