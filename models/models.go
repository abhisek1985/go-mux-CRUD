package models

type Merchant struct {
	ID int `json:"id"`
	Code string `json:"code"`
    Name string `json:"name"`
}

type Team struct {
	ID int `json:"id"`
	Email string `json:"email"`
	MerchantID *int `json:"merchant_id"`
}