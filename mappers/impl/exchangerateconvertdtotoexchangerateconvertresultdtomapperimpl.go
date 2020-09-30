package impl

import (
	"github.com/epnushtaeva/exchange_rate_converter_backend/database/entities"
	"github.com/epnushtaeva/exchange_rate_converter_backend/dto"
)

type ExchangeRateConvertDtoToExchangeRateConvertResultDtoMapperImpl struct {
}

func (exchangeRateConvertDtoToExchangeRateConvertResultDtoMapperImpl ExchangeRateConvertDtoToExchangeRateConvertResultDtoMapperImpl) ToExchangeRateConvertResultDto(
	exchangeRateConvertDto dto.ExchangeRateConvertDto,
	exchangeRate entities.ExchangeRate) dto.ExchangeRateConvertResultDto {
	exchangeRateConvertResultDto := dto.ExchangeRateConvertResultDto{}
	exchangeRateConvertResultDto.SetCurrencyFrom(exchangeRateConvertDto.GetCurrencyFrom())
	exchangeRateConvertResultDto.SetCurrencyTo(exchangeRateConvertDto.GetCurrencyTo())
	exchangeRateConvertResultDto.SetValue(exchangeRateConvertDto.GetValue())
	exchangeRateConvertResultDto.SetCourse(exchangeRate.GetCourse())
	exchangeRateConvertResultDto.SetResult(exchangeRate.GetCourse() * exchangeRateConvertDto.GetValue())
	return exchangeRateConvertResultDto
}
