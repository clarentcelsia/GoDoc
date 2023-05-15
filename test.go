package main

import (
	"time"

	// "time"

	// Model "simple-go/model"

	_ "github.com/denisenkom/go-mssqldb"

	// "simple-go/database/worker"

	"simple-go/functions"
	"simple-go/utils"

	"fmt"

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
	// worker.Basic()

	// r := gin.Default()
	// r.POST("/test", network.I_Account)
	// r.POST("/get", network.GetAccount)
	// r.POST("/getM", network.GetAccountManual)

	// r.POST("/try", utils.TryCatch)

	// r.GET("/redis", functions.BasicRedis)
	// r.Run(":9090")

	functions.ConcurrencyImplementationTypes()
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
