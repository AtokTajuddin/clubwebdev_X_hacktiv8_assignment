package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Structs sesuai requirement
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type Post struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
	Created string `json:"created_at"`
}

type Like struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

// Response format yang konsisten
type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

// In-memory storage
var (
	users  = []User{}
	posts  = []Post{}
	likes  = []Like{}
	nextID = 1
)

func main() {
	// Initialize dengan data sample
	initializeData()

	r := gin.Default()

	// Middleware logger
	r.Use(gin.Logger())

	// User endpoints
	r.POST("/users", createUser)
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	// Post endpoints
	r.POST("/posts", createPost)
	r.GET("/posts", getPosts)
	r.GET("/posts/:id", getPost)
	r.GET("/users/:id/posts", getUserPosts)
	r.DELETE("/posts/:id", deletePost)

	// Like endpoints
	r.POST("/likes", createLike)
	r.GET("/posts/:id/likes", getPostLikes)
	r.GET("/users/:id/likes", getUserLikes)

	fmt.Println("Social Media API Server starting on :8080")
	r.Run(":8080")
}

func initializeData() {
	// Sample users
	users = []User{
		{ID: "1", Username: "john_doe", Email: "john@example.com", Bio: "Hello world!"},
		{ID: "2", Username: "jane_smith", Email: "jane@example.com", Bio: "Love coding!"},
	}

	// Sample posts
	posts = []Post{
		{ID: "1", UserID: "1", Content: "My first post!", Created: time.Now().Format(time.RFC3339)},
		{ID: "2", UserID: "2", Content: "Learning Go is fun!", Created: time.Now().Format(time.RFC3339)},
	}

	nextID = 3
}

func generateID() string {
	id := strconv.Itoa(nextID)
	nextID++
	return id
}

// Helper function untuk cek username unik
func isUsernameUnique(username string, excludeID string) bool {
	for _, user := range users {
		if user.Username == username && user.ID != excludeID {
			return false
		}
	}
	return true
}

// Helper function untuk cek email unik
func isEmailUnique(email string, excludeID string) bool {
	for _, user := range users {
		if user.Email == email && user.ID != excludeID {
			return false
		}
	}
	return true
}

// Helper function untuk cek apakah user sudah like post
func hasUserLikedPost(userID, postID string) bool {
	for _, like := range likes {
		if like.UserID == userID && like.PostID == postID {
			return true
		}
	}
	return false
}

// User handlers
func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Validasi
	if newUser.Username == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Username is required",
		})
		return
	}

	if newUser.Email == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Email is required",
		})
		return
	}

	// Cek uniqueness
	if !isUsernameUnique(newUser.Username, "") {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Username already exists",
		})
		return
	}

	if !isEmailUnique(newUser.Email, "") {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Email already exists",
		})
		return
	}

	newUser.ID = generateID()
	users = append(users, newUser)

	c.JSON(http.StatusCreated, APIResponse{
		Message: "User created successfully",
		Data:    newUser,
		Error:   nil,
	})
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, APIResponse{
		Message: "Users retrieved successfully",
		Data:    users,
		Error:   nil,
	})
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, APIResponse{
				Message: "User retrieved successfully",
				Data:    user,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "User not found",
		Data:    nil,
		Error:   "User with ID " + id + " not found",
	})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")

	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			// Validasi
			if updatedUser.Username == "" {
				c.JSON(http.StatusBadRequest, APIResponse{
					Message: "Validation failed",
					Data:    nil,
					Error:   "Username is required",
				})
				return
			}

			if updatedUser.Email == "" {
				c.JSON(http.StatusBadRequest, APIResponse{
					Message: "Validation failed",
					Data:    nil,
					Error:   "Email is required",
				})
				return
			}

			// Cek uniqueness
			if !isUsernameUnique(updatedUser.Username, id) {
				c.JSON(http.StatusBadRequest, APIResponse{
					Message: "Validation failed",
					Data:    nil,
					Error:   "Username already exists",
				})
				return
			}

			if !isEmailUnique(updatedUser.Email, id) {
				c.JSON(http.StatusBadRequest, APIResponse{
					Message: "Validation failed",
					Data:    nil,
					Error:   "Email already exists",
				})
				return
			}

			updatedUser.ID = id
			users[i] = updatedUser

			c.JSON(http.StatusOK, APIResponse{
				Message: "User updated successfully",
				Data:    updatedUser,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "User not found",
		Data:    nil,
		Error:   "User with ID " + id + " not found",
	})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range users {
		if user.ID == id {
			// Hapus user
			users = append(users[:i], users[i+1:]...)

			// Hapus semua post dari user ini
			var filteredPosts []Post
			for _, post := range posts {
				if post.UserID != id {
					filteredPosts = append(filteredPosts, post)
				}
			}
			posts = filteredPosts

			// Hapus semua like dari user ini
			var filteredLikes []Like
			for _, like := range likes {
				if like.UserID != id {
					filteredLikes = append(filteredLikes, like)
				}
			}
			likes = filteredLikes

			c.JSON(http.StatusOK, APIResponse{
				Message: "User deleted successfully",
				Data:    nil,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "User not found",
		Data:    nil,
		Error:   "User with ID " + id + " not found",
	})
}

// Post handlers
func createPost(c *gin.Context) {
	var newPost Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Validasi
	if newPost.Content == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Content is required",
		})
		return
	}

	// Cek apakah user ada
	userExists := false
	for _, user := range users {
		if user.ID == newPost.UserID {
			userExists = true
			break
		}
	}

	if !userExists {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "User ID not found",
		})
		return
	}

	newPost.ID = generateID()
	newPost.Created = time.Now().Format(time.RFC3339)
	posts = append(posts, newPost)

	c.JSON(http.StatusCreated, APIResponse{
		Message: "Post created successfully",
		Data:    newPost,
		Error:   nil,
	})
}

func getPosts(c *gin.Context) {
	c.JSON(http.StatusOK, APIResponse{
		Message: "Posts retrieved successfully",
		Data:    posts,
		Error:   nil,
	})
}

func getPost(c *gin.Context) {
	id := c.Param("id")

	for _, post := range posts {
		if post.ID == id {
			c.JSON(http.StatusOK, APIResponse{
				Message: "Post retrieved successfully",
				Data:    post,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "Post not found",
		Data:    nil,
		Error:   "Post with ID " + id + " not found",
	})
}

func getUserPosts(c *gin.Context) {
	userID := c.Param("id")

	// Cek apakah user ada
	userExists := false
	for _, user := range users {
		if user.ID == userID {
			userExists = true
			break
		}
	}

	if !userExists {
		c.JSON(http.StatusNotFound, APIResponse{
			Message: "User not found",
			Data:    nil,
			Error:   "User with ID " + userID + " not found",
		})
		return
	}

	var userPosts []Post
	for _, post := range posts {
		if post.UserID == userID {
			userPosts = append(userPosts, post)
		}
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: "User posts retrieved successfully",
		Data:    userPosts,
		Error:   nil,
	})
}

func deletePost(c *gin.Context) {
	id := c.Param("id")

	for i, post := range posts {
		if post.ID == id {
			// Hapus post
			posts = append(posts[:i], posts[i+1:]...)

			// Hapus semua like untuk post ini
			var filteredLikes []Like
			for _, like := range likes {
				if like.PostID != id {
					filteredLikes = append(filteredLikes, like)
				}
			}
			likes = filteredLikes

			c.JSON(http.StatusOK, APIResponse{
				Message: "Post deleted successfully",
				Data:    nil,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "Post not found",
		Data:    nil,
		Error:   "Post with ID " + id + " not found",
	})
}

// Like handlers
func createLike(c *gin.Context) {
	var newLike Like
	if err := c.ShouldBindJSON(&newLike); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Cek apakah user ada
	userExists := false
	for _, user := range users {
		if user.ID == newLike.UserID {
			userExists = true
			break
		}
	}

	if !userExists {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "User ID not found",
		})
		return
	}

	// Cek apakah post ada
	postExists := false
	for _, post := range posts {
		if post.ID == newLike.PostID {
			postExists = true
			break
		}
	}

	if !postExists {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Post ID not found",
		})
		return
	}

	// Cek apakah user sudah like post ini
	if hasUserLikedPost(newLike.UserID, newLike.PostID) {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "User already liked this post",
		})
		return
	}

	newLike.ID = generateID()
	likes = append(likes, newLike)

	c.JSON(http.StatusCreated, APIResponse{
		Message: "Like created successfully",
		Data:    newLike,
		Error:   nil,
	})
}

func getPostLikes(c *gin.Context) {
	postID := c.Param("id")

	// Cek apakah post ada
	postExists := false
	for _, post := range posts {
		if post.ID == postID {
			postExists = true
			break
		}
	}

	if !postExists {
		c.JSON(http.StatusNotFound, APIResponse{
			Message: "Post not found",
			Data:    nil,
			Error:   "Post with ID " + postID + " not found",
		})
		return
	}

	var postLikes []Like
	for _, like := range likes {
		if like.PostID == postID {
			postLikes = append(postLikes, like)
		}
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: "Post likes retrieved successfully",
		Data:    postLikes,
		Error:   nil,
	})
}

func getUserLikes(c *gin.Context) {
	userID := c.Param("id")

	// Cek apakah user ada
	userExists := false
	for _, user := range users {
		if user.ID == userID {
			userExists = true
			break
		}
	}

	if !userExists {
		c.JSON(http.StatusNotFound, APIResponse{
			Message: "User not found",
			Data:    nil,
			Error:   "User with ID " + userID + " not found",
		})
		return
	}

	var userLikes []Like
	for _, like := range likes {
		if like.UserID == userID {
			userLikes = append(userLikes, like)
		}
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: "User likes retrieved successfully",
		Data:    userLikes,
		Error:   nil,
	})
}
