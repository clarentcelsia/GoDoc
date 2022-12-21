package main

import (
	"fmt"

	// "time"

	// Model "simple-go/model"
	// "simple-go/utils"
	"simple-go/functions"
	// "strings"
)

func init() {
	fmt.Println("this is init")
}

func main() {
	// m := new(utils.MyString)
	// isExist := m.IsExist(strList, str)
	// fmt.Println(isExist) // >>> true

	// r := gin.Default()
	// r.GET("/pagination", utils.GetFiles)
	// r.Run(":9000")

	// functions.Goroutines()
	functions.Reflect()
}
