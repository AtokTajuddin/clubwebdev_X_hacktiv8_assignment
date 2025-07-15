package main

import (
	"github.com/gin-gonic/gin"
)

type Like struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func main() {
	r := gin.Default()

	r.POST("/likes", func(c *gin.Context) {
		var newLike Like
		if err := c.ShouldBindJSON(&newLike); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		newLike.ID = "1" // Simulating ID assignment
		c.JSON(201, newLike)
	})

	r.GET("/likes", func(c *gin.Context) {
		likes := []Like{
			{ID: "1", UserID: "1", PostID: "1"},
			{ID: "2", UserID: "2", PostID: "2"},
		}
		c.JSON(200, likes)
	})

	r.GET("/likes/:id", func(c *gin.Context) {
		id := c.Param("id")
		like := Like{ID: id, UserID: "1", PostID: "1"}
		c.JSON(200, like)
	})

	r.Run()
}
