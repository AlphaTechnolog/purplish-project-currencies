package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Currency struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CompanyCurrency struct {
	CompanyID    string `json:"company_id"`
	CurrencyID   string `json:"currency_id"`
	CurrencyName string `json:"currency_name"`
	ExchangeRate int    `json:"exchange_rate"`
}

type CreateCurrencyPayload struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CreateCompanyCurrencyPayload struct {
	CompanyID    string `json:"company_id"`
	CurrencyID   string `json:"currency_id"`
	ExchangeRate int    `json:"exchange_rate"`
}

func GetCurrencies(d *sql.DB) ([]Currency, error) {
	var currencies []Currency

	rows, err := d.Query("SELECT id, name, description FROM currencies;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var currency Currency
		err = rows.Scan(&currency.ID, &currency.Name, &currency.Description)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return currencies, nil
}

func GetCurrency(d *sql.DB, ID string) (Currency, error) {
	var currency Currency

	sql := "SELECT id, name, description FROM currencies WHERE id = ? LIMIT 1;"
	row := d.QueryRow(sql, ID)
	err := row.Scan(&currency.ID, &currency.Name, &currency.Description)

	if err != nil {
		return Currency{}, err
	}

	return currency, err
}

func CreateCurrency(d *sql.DB, createPayload CreateCurrencyPayload) error {
	sql := `
        INSERT INTO currencies (id, name, description)
        VALUES
            (?, ?, ?);
    `

	_, err := d.Exec(sql, uuid.New().String(), createPayload.Name, createPayload.Description)
	if err != nil {
		return fmt.Errorf("Unable to create new currency: %w", err)
	}

	return nil
}

func RemoveCurrency(d *sql.DB, currencyID string) error {
	sql := `
        DELETE FROM currencies WHERE id = ?;
    `

	if _, err := d.Exec(sql, currencyID); err != nil {
		return fmt.Errorf("Unable to remove non-matching currency by id = '%s': %w", currencyID, err)
	}

	return nil
}

func GetCompanyCurrencies(d *sql.DB, companyID string) ([]CompanyCurrency, error) {
	var relationships []CompanyCurrency

	sql := `
		SELECT cc.company_id, cc.currency_id, c.name, cc.exchange_rate
		FROM currency_companies cc
		INNER JOIN currencies c
		ON c.id = cc.currency_id
		WHERE cc.company_id = ?;
	`

	rows, err := d.Query(sql, companyID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var element CompanyCurrency
		err = rows.Scan(&element.CompanyID, &element.CurrencyID, &element.CurrencyName, &element.ExchangeRate)
		if err != nil {
			return nil, err
		}

		relationships = append(relationships, element)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return relationships, err
}
