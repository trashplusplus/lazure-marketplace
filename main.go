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
	r.GET("/wallet/:walletId", GetProductsByWalletIdHandler(db))
	r.GET("/categories", GetAllCategoriesHandler(db))
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

		log.Print("Grabbed id from token: ", id)

		if id == -1 {
			c.IndentedJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		c.BindJSON(&product)
		log.Print("Product Price: ", product.Price)

		if product.Price <= 0 {
			c.IndentedJSON(400, gin.H{"message": "Invalid price"})
			return
		}

		AddProduct(db, product)
		c.IndentedJSON(200, gin.H{"message": "Product added successfully"})
	}
}

func GetProductByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		idFromToken := GetIdByTokenClaim(c)
		log.Print("Grabbed id from token: ", idFromToken)

		log.Print("ID: ", idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Error: ", err)
		}
		product, err := GetProductById(db, id)

		if err != nil {
			log.Println("Error: ", err)
			c.JSON(404, gin.H{"message": "Product not found"})
			return
		}

		transaction := GetTransaction(db, idStr, GrabToken(c))
		log.Println("Transaction: ", transaction.BuyerID)

		//сначала проверяем куплен ли товар, если не куплен, то проверяем овнера, если нет то удаляем ссылку
		if transaction.BuyerID != idFromToken {
			if idFromToken != product.User_Id || idFromToken == -1 {
				product.Resource_Link = ""
			}
		}

		c.IndentedJSON(200, product)
	}
}

func GetProductsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		claimId := GetIdByTokenClaim(c)
		//title param
		title := c.Query("title")

		priceStr := c.Query("price")
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			log.Println("Price error: ", err)
		}

		//category param
		categoryListStrings := c.QueryArray("category_id")
		var categoryListIds []int

		for _, strID := range categoryListStrings {
			id, err := strconv.Atoi(strID)
			if err != nil {
				log.Println("Error: ", err)
			}
			categoryListIds = append(categoryListIds, id)
		}

		if len(categoryListIds) == 0 {
			fmt.Println("Filtering products by category IDs:", categoryListIds)
		}

		for i, v := range categoryListIds {
			fmt.Println(i, v)
		}

		//next N products param
		offsetStr := c.Query("offset")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			log.Println("Offset error: ", err)
		}
		//limit param
		limitStr := c.Query("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			log.Println("Limit error: ", err)
		}

		products, err := GetProducts(db, limit, offset, title, categoryListIds, price, claimId)
		if err != nil {
			log.Println("Error: ", err)
			c.JSON(200, []Product{})
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
			log.Println("Error: ", err)
			c.JSON(200, []Product{})
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
