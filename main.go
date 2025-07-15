package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Structs sesuai requirement
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	SourceID    string  `json:"source_id"`
}

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Transaction struct {
	ID        string  `json:"id"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}

// Response format yang konsisten
type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

// In-memory storage
var (
	products     = []Product{}
	sources      = []Source{}
	transactions = []Transaction{}
	nextID       = 1
)

func main() {
	// Initialize dengan data sample
	initializeData()

	r := gin.Default()

	// Middleware logger
	r.Use(gin.Logger())

	// Product endpoints
	r.GET("/products", getProducts)
	r.GET("/products/:id", getProduct)
	r.POST("/products", createProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	// Source endpoints
	r.GET("/sources", getSources)
	r.GET("/sources/:id", getSource)
	r.POST("/sources", createSource)
	r.PUT("/sources/:id", updateSource)
	r.DELETE("/sources/:id", deleteSource)

	// Transaction endpoints
	r.POST("/transactions", createTransaction)
	r.GET("/transactions", getTransactions)
	r.GET("/transactions/:id", getTransaction)

	fmt.Println("Server starting on :8080")
	r.Run(":8080")
}

func initializeData() {
	// Sample sources
	sources = []Source{
		{ID: "1", Name: "Supplier A"},
		{ID: "2", Name: "Supplier B"},
	}

	// Sample products
	products = []Product{
		{ID: "1", Name: "Laptop", Description: "Gaming laptop", Price: 15000000, Stock: 10, SourceID: "1"},
		{ID: "2", Name: "Mouse", Description: "Wireless mouse", Price: 250000, Stock: 50, SourceID: "2"},
	}

	nextID = 3
}

func generateID() string {
	id := strconv.Itoa(nextID)
	nextID++
	return id
}

// Product handlers
func getProducts(c *gin.Context) {
	sourceID := c.Query("source_id")

	var filteredProducts []Product
	if sourceID != "" {
		for _, product := range products {
			if product.SourceID == sourceID {
				filteredProducts = append(filteredProducts, product)
			}
		}
	} else {
		filteredProducts = products
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: "Products retrieved successfully",
		Data:    filteredProducts,
		Error:   nil,
	})
}

func getProduct(c *gin.Context) {
	id := c.Param("id")

	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, APIResponse{
				Message: "Product retrieved successfully",
				Data:    product,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "Product not found",
		Data:    nil,
		Error:   "Product with ID " + id + " not found",
	})
}

func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Validasi
	if newProduct.Name == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Name is required",
		})
		return
	}

	if newProduct.Price <= 0 {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Price must be greater than 0",
		})
		return
	}

	if newProduct.Stock < 0 {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Stock must be greater than or equal to 0",
		})
		return
	}

	// Cek apakah source ada
	sourceExists := false
	for _, source := range sources {
		if source.ID == newProduct.SourceID {
			sourceExists = true
			break
		}
	}

	if !sourceExists {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Source ID not found",
		})
		return
	}

	newProduct.ID = generateID()
	products = append(products, newProduct)

	c.JSON(http.StatusCreated, APIResponse{
		Message: "Product created successfully",
		Data:    newProduct,
		Error:   nil,
	})
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")

	var updatedProduct Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	for i, product := range products {
		if product.ID == id {
			// Validasi
			if updatedProduct.Name == "" {
				c.JSON(http.StatusBadRequest, APIResponse{
					Message: "Validation failed",
					Data:    nil,
					Error:   "Name is required",
				})
				return
			}

			if updatedProduct.Price <= 0 {
				c.JSON(http.StatusBadRequest, APIResponse{
					Message: "Validation failed",
					Data:    nil,
					Error:   "Price must be greater than 0",
				})
				return
			}

			if updatedProduct.Stock < 0 {
				c.JSON(http.StatusBadRequest, APIResponse{
					Message: "Validation failed",
					Data:    nil,
					Error:   "Stock must be greater than or equal to 0",
				})
				return
			}

			updatedProduct.ID = id
			products[i] = updatedProduct

			c.JSON(http.StatusOK, APIResponse{
				Message: "Product updated successfully",
				Data:    updatedProduct,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "Product not found",
		Data:    nil,
		Error:   "Product with ID " + id + " not found",
	})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, APIResponse{
				Message: "Product deleted successfully",
				Data:    nil,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "Product not found",
		Data:    nil,
		Error:   "Product with ID " + id + " not found",
	})
}

// Source handlers
func getSources(c *gin.Context) {
	c.JSON(http.StatusOK, APIResponse{
		Message: "Sources retrieved successfully",
		Data:    sources,
		Error:   nil,
	})
}

func getSource(c *gin.Context) {
	id := c.Param("id")

	for _, source := range sources {
		if source.ID == id {
			c.JSON(http.StatusOK, APIResponse{
				Message: "Source retrieved successfully",
				Data:    source,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "Source not found",
		Data:    nil,
		Error:   "Source with ID " + id + " not found",
	})
}

func createSource(c *gin.Context) {
	var newSource Source
	if err := c.ShouldBindJSON(&newSource); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Validasi
	if newSource.Name == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Name is required",
		})
		return
	}

	newSource.ID = generateID()
	sources = append(sources, newSource)

	c.JSON(http.StatusCreated, APIResponse{
		Message: "Source created successfully",
		Data:    newSource,
		Error:   nil,
	})
}

func updateSource(c *gin.Context) {
	id := c.Param("id")

	var updatedSource Source
	if err := c.ShouldBindJSON(&updatedSource); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	for i, source := range sources {
		if source.ID == id {
			// Validasi
			if updatedSource.Name == "" {
				c.JSON(http.StatusBadRequest, APIResponse{
					Message: "Validation failed",
					Data:    nil,
					Error:   "Name is required",
				})
				return
			}

			updatedSource.ID = id
			sources[i] = updatedSource

			c.JSON(http.StatusOK, APIResponse{
				Message: "Source updated successfully",
				Data:    updatedSource,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "Source not found",
		Data:    nil,
		Error:   "Source with ID " + id + " not found",
	})
}

func deleteSource(c *gin.Context) {
	id := c.Param("id")

	for i, source := range sources {
		if source.ID == id {
			sources = append(sources[:i], sources[i+1:]...)
			c.JSON(http.StatusOK, APIResponse{
				Message: "Source deleted successfully",
				Data:    nil,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "Source not found",
		Data:    nil,
		Error:   "Source with ID " + id + " not found",
	})
}

// Transaction handlers
func createTransaction(c *gin.Context) {
	var newTransaction Transaction
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Validasi quantity
	if newTransaction.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Validation failed",
			Data:    nil,
			Error:   "Quantity must be greater than 0",
		})
		return
	}

	// Cari produk
	var product *Product
	var productIndex int
	for i, p := range products {
		if p.ID == newTransaction.ProductID {
			product = &p
			productIndex = i
			break
		}
	}

	if product == nil {
		c.JSON(http.StatusNotFound, APIResponse{
			Message: "Product not found",
			Data:    nil,
			Error:   "Product with ID " + newTransaction.ProductID + " not found",
		})
		return
	}

	// Cek stock
	if product.Stock < newTransaction.Quantity {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "Insufficient stock",
			Data:    nil,
			Error:   fmt.Sprintf("Available stock: %d, requested: %d", product.Stock, newTransaction.Quantity),
		})
		return
	}

	// Buat transaksi
	newTransaction.ID = generateID()
	newTransaction.Total = product.Price * float64(newTransaction.Quantity)

	// Kurangi stock
	products[productIndex].Stock -= newTransaction.Quantity

	// Simpan transaksi
	transactions = append(transactions, newTransaction)

	c.JSON(http.StatusCreated, APIResponse{
		Message: "Transaction created successfully",
		Data:    newTransaction,
		Error:   nil,
	})
}

func getTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, APIResponse{
		Message: "Transactions retrieved successfully",
		Data:    transactions,
		Error:   nil,
	})
}

func getTransaction(c *gin.Context) {
	id := c.Param("id")

	for _, transaction := range transactions {
		if transaction.ID == id {
			c.JSON(http.StatusOK, APIResponse{
				Message: "Transaction retrieved successfully",
				Data:    transaction,
				Error:   nil,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, APIResponse{
		Message: "Transaction not found",
		Data:    nil,
		Error:   "Transaction with ID " + id + " not found",
	})
}
