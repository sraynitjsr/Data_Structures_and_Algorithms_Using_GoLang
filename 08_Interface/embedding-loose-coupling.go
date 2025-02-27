package interfaces

import "fmt"

type Engine interface {
	Start()
}

type GasEngine struct {
	power int
}

func (e *GasEngine) Start() {
	fmt.Println("Gas engine started with power:", e.power)
}

type ElectricEngine struct {
	power int
}

func (e *ElectricEngine) Start() {
	fmt.Println("Electric engine started with power:", e.power)
}

type Car struct {
	engine Engine
	brand  string
}

func NewCar(brand string, engine Engine) *Car {
	return &Car{
		engine: engine,
		brand:  brand,
	}
}

func Embedding() {
	gasEngine := &GasEngine{power: 200}
	electricEngine := &ElectricEngine{power: 150}

	myCar := NewCar("Toyota", gasEngine)
	fmt.Println("Car brand:", myCar.brand)
	myCar.engine.Start()

	electricCar := NewCar("Tesla", electricEngine)
	fmt.Println("Car brand:", electricCar.brand)
	electricCar.engine.Start()
}
