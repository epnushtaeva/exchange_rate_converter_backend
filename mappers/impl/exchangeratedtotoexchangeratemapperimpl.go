package impl

import (
	"github.com/epnushtaeva/exchange_rate_converter_backend/database/entities"
	"github.com/epnushtaeva/exchange_rate_converter_backend/dto"
	"time"
)

type ExchangeRateDtoToExchangeRateMapperImpl struct {
}

func (exchangeRateDtoToExchangeRateMapperImpl ExchangeRateDtoToExchangeRateMapperImpl) ToExchangeRate(exchangeRateDto dto.ExchangeRateAddInDatabaseDto) entities.ExchangeRate {
	exchangeRate := entities.ExchangeRate{}
	exchangeRate.SetCurrencyFrom(exchangeRateDto.GetCurrency1())
	exchangeRate.SetCurrencyTo(exchangeRateDto.GetCurrency2())
	exchangeRate.SetUpdateDate(time.Now())
	exchangeRate.SetCourse(0)
	return exchangeRate
}
