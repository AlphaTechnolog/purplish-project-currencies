package http

import (
	"fmt"
	"net/http"

	"github.com/alphatechnolog/purplish-currencies/internal/domain"
	"github.com/alphatechnolog/purplish-currencies/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CurrencyHandler struct {
	currencyUsecase *usecase.CurrencyUsecase
}

func NewCurrencyHandler(currencyUsecase *usecase.CurrencyUsecase) *CurrencyHandler {
	return &CurrencyHandler{currencyUsecase}
}

func (h *CurrencyHandler) GetCurrencies(c *gin.Context) {
	currencies, err := h.currencyUsecase.GetCurrencies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"currencies": currencies})
}

func (h *CurrencyHandler) GetCompanyCurrencies(c *gin.Context) {
	companyID := c.Param("CompanyID")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Specify company id"})
		return
	}

	if err := uuid.Validate(companyID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID does not seem to be a valid ID"})
		return
	}

	currencies, err := h.currencyUsecase.GetCompanyCurrencies(companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"currencies": currencies})
}

func (h *CurrencyHandler) GetCurrency(c *gin.Context) {
	currencyID := c.Param("ID")
	if currencyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Specify currency id"})
		return
	}

	if err := uuid.Validate(currencyID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Currency ID does not seem to be a valid ID"})
		return
	}

	currency, err := h.currencyUsecase.GetCurrency(currencyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"currency": currency})
}

func (h *CurrencyHandler) CreateCurrency(c *gin.Context) {
	var currency domain.Currency

	if err := c.ShouldBind(&currency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request: %s", err.Error())})
		return
	}

	if err := h.currencyUsecase.CreateCurrency(&currency); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ok": true})
}

func (h *CurrencyHandler) RemoveCurrency(c *gin.Context) {
	currencyID := c.Param("ID")
	if currencyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Specify currency ID"})
		return
	}

	if err := uuid.Validate(currencyID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Currency ID does not seem to be a valid ID"})
		return
	}

	if err := h.currencyUsecase.RemoveCurrency(currencyID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
