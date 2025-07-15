package main

import (
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	SourceID    string  `json:"source_id"`
}

func main() {
	r := gin.Default()
	r.GET("/products", func(c *gin.Context) {
		products := []Product{
			{ID: "1", Name: "Product 1", Description: "Description 1", Price: 10.0, Stock: 100, SourceID: "source1"},
			{ID: "2", Name: "Product 2", Description: "Description 2", Price: 20.0, Stock: 200, SourceID: "source2"},
		}
		c.JSON(200, products)
	})
	r.Run()

	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		product := Product{ID: id, Name: "Product " + id, Description: "Description " + id, Price: 10.0 * float64(id[0]-'0'), Stock: 100, SourceID: "source" + id}
		c.JSON(200, product)
	})

	r.POST("products", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		newProduct.ID = "3" // Simulating ID assignment
		c.JSON(201, newProduct)
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedProduct Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		updatedProduct.ID = id // Simulating ID assignment
		c.JSON(200, updatedProduct)
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"message": "Product " + id + " deleted"})
	})
}
