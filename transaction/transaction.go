package main

import (
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID        string  `json:"id"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}

func main() {
	r := gin.Default()
	r.POST("/transactions", func(c *gin.Context) {
		var newTransaction Transaction
		if err := c.ShouldBindJSON(&newTransaction); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		newTransaction.ID = "1"                                        // Simulating ID assignment
		newTransaction.Total = float64(newTransaction.Quantity) * 10.0 // Simulating total calculation
		c.JSON(201, newTransaction)
	})

	r.GET("/transactions", func(c *gin.Context) {
		transactions := []Transaction{
			{ID: "1", ProductID: "product1", Quantity: 2, Total: 20.0},
			{ID: "2", ProductID: "product2", Quantity: 1, Total: 10.0},
		}
		c.JSON(200, transactions)
	})
	r.GET("/transactions/:id", func(c *gin.Context) {
		id := c.Param("id")
		transaction := Transaction{ID: id, ProductID: "product" + id, Quantity: 1, Total: 10.0}
		c.JSON(200, transaction)
	})

}
