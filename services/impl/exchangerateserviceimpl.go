package impl

import (
	"github.com/epnushtaeva/exchange_rate_converter_backend/database"
	"github.com/epnushtaeva/exchange_rate_converter_backend/database/entities"
	"github.com/epnushtaeva/exchange_rate_converter_backend/dto"
	"github.com/epnushtaeva/exchange_rate_converter_backend/mappers"
	"github.com/epnushtaeva/exchange_rate_converter_backend/services"
	"time"
)

type ExchangeRateServiceImpl struct {
	dataBaseFactory                                  database.DataBaseFactory
	exchangeRateDtoToExchangeRateMapper              mappers.ExchangeRateDtoToExchangeRateMapper
	exchangeRateConvertDtoToExchangeRateResultMapper mappers.ExchangeRateConvertDtoToExchangeRateConvertResultDtoMapper
	exchangeRateApiService                           services.ExchangeRateApiService
}

func (exchangeRateServiceImpl *ExchangeRateServiceImpl) SetDataBaseFactory(db database.DataBaseFactory) {
	exchangeRateServiceImpl.dataBaseFactory = db
}

func (exchangeRateServiceImpl *ExchangeRateServiceImpl) SetExchangeRateDtoToExchangeRateMapper(mapper mappers.ExchangeRateDtoToExchangeRateMapper) {
	exchangeRateServiceImpl.exchangeRateDtoToExchangeRateMapper = mapper
}

func (exchangeRateServiceImpl *ExchangeRateServiceImpl) SetExchangeRateConvertDtoToExchangeRateResultMapper(
	exchangeRateConvertDtoToExchangeRateResultMapper mappers.ExchangeRateConvertDtoToExchangeRateConvertResultDtoMapper) {
	exchangeRateServiceImpl.exchangeRateConvertDtoToExchangeRateResultMapper = exchangeRateConvertDtoToExchangeRateResultMapper
}

func (exchangeRateServiceImpl *ExchangeRateServiceImpl) SetExchangeRateApiService(exchangeRateApiService services.ExchangeRateApiService) {
	exchangeRateServiceImpl.exchangeRateApiService = exchangeRateApiService
}

func (exchangeRateServiceImpl ExchangeRateServiceImpl) Add(exchangeRateDto dto.ExchangeRateAddInDatabaseDto) dto.AddEntityResultDto {
	if exchangeRateServiceImpl.Exists(exchangeRateDto) {
		return exchangeRateServiceImpl.GetAddExchangeExistsErrorResult()
	}

	exchangeRate := exchangeRateServiceImpl.exchangeRateDtoToExchangeRateMapper.ToExchangeRate(exchangeRateDto)
	exchangeRate.SetCourse(exchangeRateServiceImpl.exchangeRateApiService.GetCourse(exchangeRateDto))

	exchangeRateServiceImpl.dataBaseFactory.GetDB().Create(&exchangeRate)
	return exchangeRateServiceImpl.GetAddExchangeSuccessResult()
}

func (exchangeRateServiceImpl ExchangeRateServiceImpl) Convert(exchangeRateConvertDto dto.ExchangeRateConvertDto) dto.ExchangeRateConvertResultDto {
	if exchangeRateServiceImpl.NotExists(exchangeRateConvertDto) {
		return exchangeRateServiceImpl.GetEmptyConvertResultDto(exchangeRateConvertDto)
	}

	exchangeRate := entities.ExchangeRate{}
	exchangeRateServiceImpl.
		dataBaseFactory.
		GetDB().
		Where(
			"currency_from = ? AND currency_to = ?",
			exchangeRateConvertDto.GetCurrencyFrom(),
			exchangeRateConvertDto.GetCurrencyTo()).
		Find(&exchangeRate)

	return exchangeRateServiceImpl.exchangeRateConvertDtoToExchangeRateResultMapper.ToExchangeRateConvertResultDto(
		exchangeRateConvertDto,
		exchangeRate)
}

func (exchangeRateServiceImpl ExchangeRateServiceImpl) UpdateAll(updateTime time.Time) {
	var exchangeRates []entities.ExchangeRate
	exchangeRateServiceImpl.dataBaseFactory.GetDB().Where("course=0").Find(&exchangeRates)

	for _, exchangeRate := range exchangeRates {
		exchangeRate.SetCourse(exchangeRateServiceImpl.exchangeRateApiService.GetCourse(dto.ExchangeRateAddInDatabaseDto{
			Currency1: exchangeRate.GetCurrencyFrom(),
			Currency2: exchangeRate.GetCurrencyTo()}))
		exchangeRate.SetUpdateDate(updateTime)
		exchangeRateServiceImpl.dataBaseFactory.GetDB().Save(&exchangeRate)
	}
}

func (exchangeRateServiceImpl ExchangeRateServiceImpl) Exists(exchangeRate dto.ExchangeRateAddInDatabaseDto) bool {
	result := exchangeRateServiceImpl.
		dataBaseFactory.
		GetDB().
		Where(
			"currency_from = ? AND currency_to = ?",
			exchangeRate.GetCurrency1(),
			exchangeRate.GetCurrency2()).
		First(&entities.ExchangeRate{})

	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func (exchangeRateServiceImpl ExchangeRateServiceImpl) NotExists(exchangeRate dto.ExchangeRateConvertDto) bool {
	result := exchangeRateServiceImpl.
		dataBaseFactory.
		GetDB().
		Where(
			"currency_from = ? AND currency_to = ?",
			exchangeRate.GetCurrencyFrom(),
			exchangeRate.GetCurrencyTo()).
		First(&entities.ExchangeRate{})

	if result.RowsAffected == 0 {
		return true
	} else {
		return false
	}
}

func (exchangeRateServiceImpl ExchangeRateServiceImpl) GetAddExchangeExistsErrorResult() dto.AddEntityResultDto {
	return dto.AddEntityResultDto{
		Status:  "error",
		Message: "Запись с данными значениями валют уже существует в базе данных"}
}

func (exchangeRateServiceImpl ExchangeRateServiceImpl) GetAddExchangeSuccessResult() dto.AddEntityResultDto {
	return dto.AddEntityResultDto{
		Status:  "success",
		Message: "Запись успешно добавлена"}
}

func (exchangeRateServiceImpl ExchangeRateServiceImpl) GetEmptyConvertResultDto(exchangeRateConvertDto dto.ExchangeRateConvertDto) dto.ExchangeRateConvertResultDto {
	result := dto.ExchangeRateConvertResultDto{}
	result.SetCurrencyFrom(exchangeRateConvertDto.GetCurrencyFrom())
	result.SetCurrencyTo(exchangeRateConvertDto.GetCurrencyTo())
	result.SetValue(exchangeRateConvertDto.GetValue())
	result.SetCourse(0)
	result.SetResult(0)
	return result
}
