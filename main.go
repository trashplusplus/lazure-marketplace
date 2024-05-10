package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := InitDB()
	if err != nil {
		log.Println("Error initializing database:", err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set")
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{"message": "productAPI of lazure-marketplace"})
	})

	r.POST("/product", AddProductHandler(db))
	r.GET("/product/:id", GetProductByIdHandler(db))
	r.DELETE("/product/:id", DeleteProductByIdHandler(db))
	r.GET("/catalog", GetProductsByTitleHandler(db))
	r.GET("/wallet/:walletId", GetProductsByWalletIdHandler(db))
	r.GET("/category", GetAllCategoriesHandler(db))
	r.GET("/get-products", GetProductsHandler(db))

	serverAddress := ":" + port
	log.Printf("Starting server on %s...", serverAddress)

	if err := r.Run(serverAddress); err != nil {
		log.Fatal("Failed to start server:", err)
	}

	defer db.Close()
}

func AddProductHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product Product
		id := GetIdByTokenClaim(c)
		product.User_Id = id

		if id == -1 {
			return
		}

		c.BindJSON(&product)
		AddProduct(db, product)
		c.IndentedJSON(200, gin.H{"message": "Product added successfully"})
	}
}

func GetProductByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		log.Print("ID: ", idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Error: ", err)
		}
		product, err := GetProductById(db, id)
		if err != nil {
			log.Println("Error: ", err)
			c.JSON(404, gin.H{"message": "Product not found"})
		} else {
			c.IndentedJSON(200, product)
		}

	}
}

// GetProductByTitleHandler returns a Gin handler function for retrieving a product by its title.
//
// It takes a *sql.DB as a parameter and returns a gin.HandlerFunc.
func GetProductsByTitleHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(404, gin.H{"message": "Search is empty"})
			return
		}

		products, err := GetProductsByTitle(db, name)
		if err != nil {
			log.Println("Error: ", err)
			c.JSON(404, gin.H{"message": "No products by that name was found"})
		} else {
			c.IndentedJSON(200, products)
		}
	}
}

func GetProductsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		defaultLimit := 20
		claimId := GetIdByTokenClaim(c)
		limit := c.Query("limit")
		limit_int, err := strconv.Atoi(limit)
		if err != nil {
			log.Println("Limit error: ", err)
		}
		defaultLimit = limit_int

		products, err := GetProducts(db, defaultLimit, claimId)
		if err != nil {
			log.Println("Error: ", err)
			c.JSON(404, gin.H{"message": "No products were found"})
		} else {
			c.IndentedJSON(200, products)
		}
	}
}

func GetProductsByWalletIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//-1 means unuthorized
		id := GetIdByTokenClaim(c)

		walletId := c.Param("walletId")

		log.Print("Grabbed id from token: ", id)

		products, err := GetProductsByWalletId(db, walletId, id)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}

		if len(products) == 0 {
			c.JSON(200, []Product{})
			return
		}
		c.IndentedJSON(200, products)

	}
}

func GetAllCategoriesHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := GetAllCategories(db)
		if err != nil {
			log.Println("Error: ", err)
			c.JSON(404, gin.H{"message": "No categories were found"})
		} else {
			c.IndentedJSON(200, categories)
		}
	}
}

func DeleteProductByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Error: ", err)
		}
		claimId := GetIdByTokenClaim(c)

		productName, err := DeleteProductById(db, id, claimId)
		if err != nil {
			c.JSON(404, gin.H{"message": "Product wasn't deleted"})
		} else {
			c.IndentedJSON(200, gin.H{"message": fmt.Sprintf("[%s] has been deleted", productName)})
		}

	}
}
