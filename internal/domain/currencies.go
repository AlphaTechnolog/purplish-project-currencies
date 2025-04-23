package domain

import "github.com/google/uuid"

type Currency struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description"`
}

type CurrencyCompany struct {
	CompanyID string `json:"company_id"`
	CurrencyID string `json:"currency_id"`
	CurrencyName string `json:"currency_name"`
	ExchangeRate int `json:"exchange_rate"`
}

func (cc *CurrencyCompany) ValidateUUIDs() error {
	if err := uuid.Validate(cc.CompanyID); err != nil {
		return err
	}
	if err := uuid.Validate(cc.CurrencyID); err != nil {
		return err
	}
	return nil
}
