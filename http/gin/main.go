package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Obj struct {
	Name string `uri:"name" binding:"required"`
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before")
		c.Next()
		log.Println("after")
	}
}

func main() {
	r := gin.Default()
	r.Use(Logger())

	r.GET("/user/:name", func(c *gin.Context) {
		var obj Obj
		if err := c.ShouldBindUri(&obj); err != nil {
			c.JSON(400, gin.H{"code": -1})
			return
		}

		c.JSON(200, gin.H{"name": obj.Name})
	})

	r.Run(":8080")
}
