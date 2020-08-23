package factory

import "log"

type Car interface {
	Run(speed int)
}

type Benz struct {
}

type Tesla struct {
}

func (car *Benz) Run(speed int) {
	log.Println("Benz", speed)
}

func (car *Tesla) Run(speed int) {
	log.Println("Tesla", speed)
}

func NewCar(t int) Car {

	if t == 1 {
		return &Tesla{}
	} else {
		return &Benz{}
	}
}
