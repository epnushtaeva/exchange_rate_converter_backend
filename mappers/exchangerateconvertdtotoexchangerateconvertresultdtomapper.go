package mappers

import (
	"github.com/epnushtaeva/exchange_rate_converter_backend/database/entities"
	"github.com/epnushtaeva/exchange_rate_converter_backend/dto"
)

type ExchangeRateConvertDtoToExchangeRateConvertResultDtoMapper interface {
	ToExchangeRateConvertResultDto(exchangeRateConvertDto dto.ExchangeRateConvertDto, exchangeRate entities.ExchangeRate) dto.ExchangeRateConvertResultDto
}
