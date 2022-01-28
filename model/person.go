package model

import (
	"time"
)

type person struct {
	name      string
	age       int
	dob       string
	createdAt time.Time
}

//"getter"
func (p person) GetName() string{
	return p.name
}

func (p person) GetAge() int{
	return p.age
}

func (p person) GetDOB() string{
	return p.dob
}

// func GetNameTest(p person) string{
// 	return.p
// }

//"setter"
func SetNewPerson(name string, age int, dob string) person{
	//call person type
	return person{
		name: name,
		age: age,
		dob: dob,
	}
}

func SetNewPersonName(name string) person{
	//call person type
	return person{
		name: name,
	}
}