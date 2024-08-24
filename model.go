package main

type Message struct {
	MMID            string `json:"MMID"`
	TransactionType string
	ProductType     string
	PaymentAmount   float32
	TransferType    string
}
