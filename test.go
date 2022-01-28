package main

import (
	"fmt"
	// "time"

	// Model "simple-go/model"
	Funct "simple-go/functions"
	// "strings"
)

func init() {
	fmt.Println("this is init")
}

func main() {
	// person := Model.SetNewPerson("Hanny", 20, "22-02-12")
	// name := person.GetName()
	// fmt.Println(name)

	//https://pkg.go.dev/strings#Contains
	// data := []string{"apple", "mango"}
	// var result []string = Funct.FunctionAsParams(data, func(cases string) bool{
	// 	//todo: what do you want to do as a feedback of the func == string[]
	// 	return strings.Contains(cases, "o") //return bool whether "o" is within cases
	// })

	// fmt.Println(result)

	Funct.Interface()
}


