package mappers

import (
	"github.com/epnushtaeva/exchange_rate_converter_backend/database/entities"
	"github.com/epnushtaeva/exchange_rate_converter_backend/dto"
)

type ExchangeRateDtoToExchangeRateMapper interface {
	ToExchangeRate(exchangeRateDto dto.ExchangeRateAddInDatabaseDto) entities.ExchangeRate
}
