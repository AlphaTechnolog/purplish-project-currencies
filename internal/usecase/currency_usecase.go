package usecase

import (
	"fmt"

	"github.com/alphatechnolog/purplish-currencies/internal/domain"
	"github.com/alphatechnolog/purplish-currencies/internal/repository"
	"github.com/google/uuid"
)

type CurrencyUsecase struct {
	sqldbRepo repository.SQLDBRepository
}

func NewCurrencyUsecase(sqldbRepo repository.SQLDBRepository) *CurrencyUsecase {
	return &CurrencyUsecase{sqldbRepo}
}

func (uc *CurrencyUsecase) GetCurrencies() ([]domain.Currency, error) {
	sql := "SELECT id, name, description FROM currencies;"
	currencies := []domain.Currency{}

	rows, err := uc.sqldbRepo.Query(sql)
	if err != nil {
		return currencies, fmt.Errorf("unable to query currencies: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var currency domain.Currency
		if err := rows.Scan(&currency.ID, &currency.Name, &currency.Description); err != nil {
			return currencies, fmt.Errorf("cannot scan currency: %w", err)
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (uc *CurrencyUsecase) GetCurrency(id string) (*domain.Currency, error) {
	sql := "SELECT id, name, description FROM currencies WHERE id = ? LIMIT 1;"
	row := uc.sqldbRepo.QueryRow(sql, id)

	currency := &domain.Currency{}
	err := row.Scan(&currency.ID, &currency.Name, &currency.Description)
	if err != nil {
		return nil, fmt.Errorf("unable to scan currency: %w", err)
	}

	return currency, nil
}

func (uc *CurrencyUsecase) CreateCurrency(currency *domain.Currency) error {
	query := "INSERT INTO currencies (id, name, description) VALUES (?, ?, ?)"
	currency.ID = uuid.New().String()
	_, err := uc.sqldbRepo.Execute(query, currency.ID, currency.Name, currency.Description)
	if err != nil {
		return fmt.Errorf("failed to create currency: %w", err)
	}

	return nil
}

func (uc *CurrencyUsecase) RemoveCurrency(id string) error {
	sql := "DELETE FROM currencies WHERE id = ?;"

	if _, err := uc.sqldbRepo.Execute(sql, id); err != nil {
		return err
	}

	return nil
}

func (uc *CurrencyUsecase) GetCompanyCurrencies(companyID string) ([]domain.CurrencyCompany, error) {
	sql := `
		SELECT cc.company_id, cc.currency_id, c.name, c.exchange_rate
		FROM currency_companies cc
		INNER JOIN currencies c
		ON c.id = cc.currency_id
		WHERE cc.company_id = ?;
	`

	rows, err := uc.sqldbRepo.Query(sql, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	relationships := []domain.CurrencyCompany{}

	for rows.Next() {
		var element domain.CurrencyCompany
		err = rows.Scan(&element.CompanyID, &element.CurrencyID, &element.CurrencyName, &element.ExchangeRate)
		if err != nil {
			return nil, err
		}

		relationships = append(relationships, element)
	}

	return relationships, nil
}
