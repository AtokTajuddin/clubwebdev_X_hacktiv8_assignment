package main

import (
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.GET("/sources", func(c *gin.Context) {
		sources := []string{"source1", "source2", "source3"}
		c.JSON(200, sources)
	})

	r.GET("/sources/:id", func(c *gin.Context) {
		id := c.Param("id")
		source := "source" + id
		c.JSON(200, gin.H{"id": id, "source": source})
	})

	r.POST("/sources", func(c *gin.Context) {
		var newSource string
		if err := c.ShouldBindJSON(&newSource); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"source": newSource})
	})
	r.PUT("/sources/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedSource string
		if err := c.ShouldBindJSON(&updatedSource); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"id": id, "source": updatedSource})
	})

	r.DELETE("/sources/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"id": id, "deleted": true})
	})
}
