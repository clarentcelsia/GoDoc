package model

type Student struct{
	Nim string
	Email string
}

func (s Student) ToString() string{
	return "Student{"+
		"Nim: " + s.Nim + " " +
		"Email: " + s.Email +
		"}"
}