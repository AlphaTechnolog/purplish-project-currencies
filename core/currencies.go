package core

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/alphatechnolog/purplish-currencies/database"
	"github.com/gin-gonic/gin"
)

func getCurrencies(d *sql.DB, c *gin.Context) error {
	currencies, err := database.GetCurrencies(d)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"currencies": currencies})

	return nil
}

func getCompanyCurrencies(d *sql.DB, c *gin.Context) error {
	companyID := c.Param("CompanyID")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Specify company ID"})
		return nil
	}

	currencies, err := database.GetCompanyCurrencies(d, companyID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"currencies": currencies})

	return nil
}

func getCurrency(d *sql.DB, c *gin.Context) error {
	currencyID := c.Param("ID")
	if currencyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Specify currency ID"})
		return nil
	}

	currency, err := database.GetCurrency(d, currencyID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"currency": currency})

	return nil
}

func createCurrency(d *sql.DB, c *gin.Context) error {
	bodyContents, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	var createPayload database.CreateCurrencyPayload
	if err = json.Unmarshal(bodyContents, &createPayload); err != nil {
		return err
	}

	if err = database.CreateCurrency(d, createPayload); err != nil {
		return err
	}

	c.JSON(http.StatusCreated, gin.H{"ok": true})

	return nil
}

func removeCurrency(d *sql.DB, c *gin.Context) error {
	currencyID := c.Param("ID")
	if currencyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return nil
	}

	if err := database.RemoveCurrency(d, currencyID); err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})

	return nil
}

func CreateCurrenciesRoutes(d *sql.DB, r *gin.RouterGroup) {
	r.GET("/", WrapError(WithDB(d, getCurrencies)))
	r.GET("/company-currencies/:CompanyID", WrapError(WithDB(d, getCompanyCurrencies)))
	r.GET("/:ID", WrapError(WithDB(d, getCurrency)))
	r.POST("/", WrapError(WithDB(d, createCurrency)))
	r.DELETE("/:ID", WrapError(WithDB(d, removeCurrency)))
}
