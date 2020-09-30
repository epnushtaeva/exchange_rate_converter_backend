package services

import (
	"github.com/epnushtaeva/exchange_rate_converter_backend/dto"
	"time"
)

type ExchangeRateService interface {
	Add(exchangeRateDto dto.ExchangeRateAddInDatabaseDto) dto.AddEntityResultDto
	Convert(exchangeRateConvertDto dto.ExchangeRateConvertDto) dto.ExchangeRateConvertResultDto
	UpdateAll(updateTime time.Time)
}
