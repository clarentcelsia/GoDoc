package main

import (
	"fmt"
	"time"

	// "time"

	// Model "simple-go/model"
	"simple-go/functions"
	"simple-go/utils"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	// "strings"
)

func init() {
	fmt.Println("this is init")
}

func Middleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(1*time.Minute),
		timeout.WithHandler(func(c *gin.Context) {
			fmt.Println("Done executing, pass on request handler to the next-in-chain")
			c.Next()
		}),
		// timeout.WithResponse(func(c *gin.Context) {
		// 	c.String(http.StatusRequestTimeout, "Server request timeout!")
		// }),
	)

}

func main() {
	// m := new(utils.MyString)
	// isExist := m.IsExist(strList, str)
	// fmt.Println(isExist) // >>> true

	// r := gin.Default()
	// r.Use(Middleware())

	// r.GET("/pagination", utils.GetFiles)
	// r.GET("/", Handler)
	// r.Run(":9000")

	functions.Mutex()
}

func Handler(c *gin.Context) {
	// ======== TIME EXEC FLAGGING ==========
	t := time.Now()
	defer fmt.Printf("TIME EXEC : %v", utils.Tracker(t))

	// This already handled by Middleware
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	// defer cancel()
	// ctx := c.Request.WithContext(context.Background()).Context()

	select {
	case <-c.Done():
		fmt.Println("Server Timeout!")
	case result := <-functions.Goroutines(c):
		c.JSON(200, result)
	}

}
