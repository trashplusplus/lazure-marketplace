package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Transaction struct {
	TransactionID int    `json:"transactionId"`
	SellerID      int    `json:"sellerId"`
	BuyerID       int    `json:"buyerId"`
	DateTime      string `json:"dateTime"`
	ProductID     int    `json:"productId"`
	TxID          string `json:"txId"`
}

func GetTransaction(db *sql.DB, idStr string, token string) Transaction {
	client := &http.Client{}
	url := "https://accountsapi-3a5f92f4b3d5.herokuapp.com/Transactions/product/" + idStr
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request to transactions API: ", err)
		return Transaction{}
	}

	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to transactions API: ", err)
		return Transaction{}
	}
	defer resp.Body.Close()

	var transactions []Transaction

	if err := json.NewDecoder(resp.Body).Decode(&transactions); err != nil {
		log.Println("Error decoding response from transactions API: ", err)
		return Transaction{}
	}

	if len(transactions) > 0 {
		for _, transaction := range transactions {
			return transaction
		}
	}

	return Transaction{}

}
