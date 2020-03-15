package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "12 factor app 0.1 GAMMA",
		})
	})

	var v1 = router.Group("/api/v1")
	{
		var productsRoot = v1.Group("products")
		{
			productsRoot.GET("/", func(c *gin.Context) {

				rows, _ := find("product")

				c.JSON(200, gin.H{
					"products": rows,
				})
			})
			productsRoot.POST("/:name", func(c *gin.Context) {
				name := c.Param("name")

				row, err := add("product", name)

				if err != nil {
					c.JSON(400, gin.H{
						"error": err.Error(),
					})

					return
					
				}

				c.JSON(201, gin.H{
					"message": fmt.Sprintf("Added a product called %s", name),
					"row":     row,
				})

			})
		}
	}
}

func start() error {
	return router.Run(":9090")
}
