package main

import (
	"github.com/gin-gonic/gin"
)

type Post struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
	Created string `json:"created_at"`
}

func main() {
	r := gin.Default()

	r.GET("/posts", func(c *gin.Context) {
		posts := []Post{
			{ID: "1", UserID: "1", Content: "This is the first post", Created: "2023-10-01T12:00:00Z"},
			{ID: "2", UserID: "2", Content: "This is the second post", Created: "2023-10-02T12:00:00Z"},
		}
		c.JSON(200, posts)
	})

	r.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		post := Post{ID: id, UserID: "1", Content: "This is a post with ID " + id, Created: "2023-10-01T12:00:00Z"}
		c.JSON(200, post)
	})

	r.POST("/posts", func(c *gin.Context) {
		var newPost Post
		if err := c.ShouldBindJSON(&newPost); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		newPost.ID = "3" // Simulating ID assignment
		c.JSON(201, newPost)
	})

	r.PUT("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedPost Post
		if err := c.ShouldBindJSON(&updatedPost); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		updatedPost.ID = id // Simulating ID assignment
		c.JSON(200, updatedPost)
	})

	r.DELETE("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"message": "Post " + id + " deleted"})
	})

	r.Run()
}
