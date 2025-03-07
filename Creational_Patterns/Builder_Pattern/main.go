package main

import "fmt"

type Car struct { //This model is the main car which will be build from the below CarBuilder model
	Brand      string
	Model      string
	Color      string
	EngineType string
}

type CarBuilder struct {
	brand      string
	model      string
	color      string
	engineType string
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (cb *CarBuilder) SetBrand(brand string) *CarBuilder {
	cb.brand = brand
	return cb
}

func (cb *CarBuilder) SetModel(model string) *CarBuilder {
	cb.model = model
	return cb
}

func (cb *CarBuilder) SetColor(color string) *CarBuilder {
	cb.color = color
	return cb
}

func (cb *CarBuilder) SetEngineType(engineType string) *CarBuilder {
	cb.engineType = engineType
	return cb
}

func (cb *CarBuilder) Build() Car {
	return Car{
		Brand:      cb.brand,
		Model:      cb.model,
		Color:      cb.color,
		EngineType: cb.engineType,
	}
}
func main() {

	car := NewCarBuilder().
		SetBrand("audi").
		SetModel("A6").
		SetColor("Black").
		SetEngineType("2000cc")

	fmt.Println("Car Detail are %v ", car)
}
