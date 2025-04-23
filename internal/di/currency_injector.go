package di

import (
	"database/sql"

	"github.com/alphatechnolog/purplish-currencies/delivery/http"
	"github.com/alphatechnolog/purplish-currencies/infrastructure/database"
	"github.com/alphatechnolog/purplish-currencies/internal/usecase"
	"github.com/gin-gonic/gin"
)

type CurrencyInjector struct {
	db *sql.DB
}

func NewCurrencyInjector(db *sql.DB) ModuleInjector {
	return &CurrencyInjector{db}
}

func (ci *CurrencyInjector) Inject(r *gin.RouterGroup) {
	sqliteRepo := database.NewSQLiteRepository(ci.db)
	currencyUsecase := usecase.NewCurrencyUsecase(sqliteRepo)
	currencyHandler := http.NewCurrencyHandler(currencyUsecase)

	r.GET("/", currencyHandler.GetCurrencies)
	r.GET("/:ID", currencyHandler.GetCurrency)
	r.GET("/company-currencies/:CompanyID", currencyHandler.GetCompanyCurrencies)
	r.POST("/", currencyHandler.CreateCurrency)
	r.DELETE("/:ID", currencyHandler.RemoveCurrency)
}
