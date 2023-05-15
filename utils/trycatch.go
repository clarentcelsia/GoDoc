package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Exception interface{}

type (
	Block struct {
		Try     func()
		Catch   func(e Exception)
		Finally func()
	}
)

// note. this panic will be run after defer executed.
func Throw(e Exception) {
	panic(e)
}

func (b *Block) Do() {
	if b.Finally != nil {
		// Check and run a statement inside this block
		// If exists, call finally() with a 'defer' to make it being executed the last.
		defer b.Finally()
	}
	if b.Catch != nil {
		// triggers if the error occurred (panic)
		defer func() {
			// recover the program, then catch the error
			if r := recover(); r != nil {
				b.Catch(r)
			}
		}()
	}
	b.Try()
}

func TryCatch(c *gin.Context) {
	var block = Block{
		Try: func() {
			x := 1
			y := x / (x - 1)
			fmt.Println(y)
		},
		Catch: func(e Exception) {
			fmt.Println("This error printed by catch func")
			c.JSON(500, gin.H{
				"errormsg": e.(interface{}), // or just e
			})
		},
		// Finally: func() {
		// 	println("Finally Block")
		// },
	}

	block.Do()
}
