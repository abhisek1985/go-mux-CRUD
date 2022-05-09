package models

import (
    "net/url"
)

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

func (a *Merchant) ValidateMerchant() url.Values {
	errs := url.Values{}

	// check if the code empty
	if a.Code == "" {
		errs.Add("code", "The code field is required!")
	}

	// check if the name empty
    if a.Name == "" {
        errs.Add("name", "The name field is required!")
    }

	return errs
}

func (a *Team) ValidateTeam() url.Values {
	errs := url.Values{}

	// check if the email empty
	if a.Email == "" {
		errs.Add("email", "The email field is required!")
	}

	return errs
}