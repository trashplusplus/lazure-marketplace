// i need to create productsAPI, a server to connect to the Postgresql db and get data from it
// basic CRUD operations and some different files, for DB connection postgresql.go and products.go
// also i need to import Gin library
package main

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	serverIP := "localhost:3030"
	db, err := InitDB()
	if err != nil {
		log.Println("Error: ", err)
	}

	r := gin.Default()
	r.POST("/add", AddProductHandler(db))
	r.GET("/getbyid/:id", GetProductByIdHandler(db))
	r.GET("/getbytitle", GetProductsByTitleHandler(db))

	r.Run(serverIP)
	defer db.Close()
}

func AddProductHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product Product
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
		title := c.Query("title")
		product, err := GetProductsByTitle(db, title)
		if err != nil {
			log.Println("Error: ", err)
			c.JSON(404, gin.H{"message": "No products by that name was found"})
		} else {
			c.IndentedJSON(200, product)
		}
	}
}
