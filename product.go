package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	Product_Id    int     `json:"product_id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	Resource_Link string  `json:"resource_link"`
	Category_Id   int     `json:"category_id"`
	User_Id       int     `json:"user_id"`
}

type ProductToBuy struct {
	Product_Id  int     `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category_Id int     `json:"category_id"`
	User_Id     int     `json:"user_id"`
}

type Category struct {
	Category_Id int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func AddProduct(db *sql.DB, product Product) error {
	//todo validating by TOKEN
	statement := `insert into Products(name, description, price, user_id, resource_link, category_id) values($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(statement, product.Name, product.Description, product.Price, product.User_Id, product.Resource_Link, product.Category_Id)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	log.Printf("Product %s added successfully", product.Name)
	return nil
}

func GetProductById(db *sql.DB, id int) (*ProductToBuy, error) {
	var product ProductToBuy
	row := db.QueryRow("SELECT product_id, name, description, price, user_id, category_id FROM Products WHERE product_id = $1", id)
	err := row.Scan(&product.Product_Id, &product.Name, &product.Description, &product.Price, &product.User_Id, &product.Category_Id)
	if err != nil {
		log.Println("GetProductById error: ", err)
		return nil, err
	}
	return &product, nil
}

func GetProductsByTitle(db *sql.DB, title string) ([]ProductToBuy, error) {
	var products []ProductToBuy
	rows, err := db.Query("Select product_id, name, description, price, user_id, category_id from Products where name ILIKE '%' || $1 || '%'", title)
	if err != nil {
		log.Println("GetProductByTitle error: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product ProductToBuy
		err := rows.Scan(&product.Product_Id, &product.Name, &product.Description, &product.Price, &product.User_Id, &product.Category_Id)
		if err != nil {
			log.Println("GetProductByTitle error: ", err)
			continue
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Println("GetProductByTitle error: ", err)
		return nil, err
	}

	if len(products) == 0 {
		return nil, sql.ErrNoRows
	}

	return products, nil
}

func GetProductsByWalletId(db *sql.DB, walletId string) ([]ProductToBuy, error) {

	var products []ProductToBuy
	rows, err := db.Query("select p.product_id, p.name, p.description, p.price, p.user_id, p.category_id from Products p join Users u on p.user_id = u.user_id where u.wallet_id = $1", walletId)
	if err != nil {
		log.Println("GetProductsByWalletId error: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product ProductToBuy
		err := rows.Scan(&product.Product_Id, &product.Name, &product.Description, &product.Price, &product.User_Id, &product.Category_Id)
		if err != nil {
			log.Println("GetProductsByWalletId error: ", err)
			continue
		}
		products = append(products, product)
	}

	if products == nil {
		return nil, sql.ErrNoRows
	}

	return products, nil
}

func GetAllCategories(db *sql.DB) ([]Category, error) {

	var categories []Category
	rows, err := db.Query("select category_id, name, description from Categories")
	if err != nil {
		log.Println("GetAllCategories error: ", err)
		return nil, err
	}
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.Category_Id, &category.Name, &category.Description)
		if err != nil {
			log.Println("GetAllCategories error: ", err)
			continue
		}
		categories = append(categories, category)
	}

	if categories == nil {
		return nil, sql.ErrNoRows
	}

	return categories, nil
}

func DeleteProductById(db *sql.DB, product_id int, user_id int) (string, error) {
	var productName string

	err := db.QueryRow("SELECT name FROM Products WHERE product_id = $1 AND user_id = $2", product_id, user_id).Scan(&productName)
	if err != nil {
		log.Println("DeleteProductById error:", err)
		return "", err
	}

	_, err = db.Exec("DELETE FROM Products WHERE product_id = $1 AND user_id = $2", product_id, user_id)
	if err != nil {
		log.Println("DeleteProductById error:", err)
		return "", err
	}

	return productName, nil
}
