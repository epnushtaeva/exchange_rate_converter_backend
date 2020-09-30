package services

import "github.com/epnushtaeva/exchange_rate_converter_backend/dto"

type ExchangeRateApiService interface {
	GetCourse(exchangeRateDto dto.ExchangeRateAddInDatabaseDto) float32
}
