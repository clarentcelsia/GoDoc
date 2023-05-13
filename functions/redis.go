package functions

import (
	"context"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
)

// Install redis server to use redis**

/* Caching
Generally Redis is used for database, cache, session management, message broker, queue
relevant to real-time operations and temporary data.*/
func BasicRedis(c *gin.Context) {
	// var host = "localhost:6379"
	// var pass = ""

	// initialize redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:9090",
		Password: "",
		DB:       0,
	})
	// check if redis has been set up
	status, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"status": status,
	})
}
