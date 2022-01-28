package packages

import (
	"time"

	"../model"
)

func NewPerson(name string, age int, dob string, createdAt time.Time) model.Person {
	return model.Person{
		Name:      name,
		Age:       age,
		DOB:       dob,
		CreatedAt: createdAt,
	}
}

func NewCar(name string, color string) *model.Car {
	return &model.Car{
		Name:  name,
		Color: color,
	}
}
