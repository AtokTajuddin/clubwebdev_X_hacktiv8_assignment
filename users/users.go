package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

func main() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		newUser.ID = "3" // Simulating ID assignment
		c.JSON(201, newUser)
	})

	r.GET("/users", func(c *gin.Context) {
		users := []User{
			{ID: "1", Username: "user1", Email: "user1@example.com", Bio: "Bio of user1"},
			{ID: "2", Username: "user2", Email: "user2@example.com", Bio: "Bio of user2"},
		}
		c.JSON(200, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := User{ID: id, Username: "user" + id, Email: "user" + id + "@example.com", Bio: "Bio of user" + id}
		c.JSON(200, user)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedUser User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		updatedUser.ID = id // Simulating ID assignment
		c.JSON(200, updatedUser)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"message": "User " + id + " deleted"})
	})

}
