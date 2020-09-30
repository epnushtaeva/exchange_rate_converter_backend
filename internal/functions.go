package internal

import (
	"github.com/epnushtaeva/exchange_rate_converter_backend/database"
	"github.com/epnushtaeva/exchange_rate_converter_backend/database/entities"
	"github.com/epnushtaeva/exchange_rate_converter_backend/dto"
	mappers "github.com/epnushtaeva/exchange_rate_converter_backend/mappers/impl"
	services_interfaces "github.com/epnushtaeva/exchange_rate_converter_backend/services"
	services "github.com/epnushtaeva/exchange_rate_converter_backend/services/impl"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func InitDataBaseFactory(connectionString string) database.DataBaseFactory {
	dataBaseFactory := database.DataBaseFactory{}
	dataBaseFactory.InitDBConnection(connectionString)
	dataBaseFactory.GetDB().AutoMigrate(&entities.ExchangeRate{})
	return dataBaseFactory
}

func InitExchangeRateService(dataBaseFactory database.DataBaseFactory) services_interfaces.ExchangeRateService {
	exchangeRateService := services.ExchangeRateServiceImpl{}
	exchangeRateService.SetDataBaseFactory(dataBaseFactory)
	exchangeRateService.SetExchangeRateDtoToExchangeRateMapper(mappers.ExchangeRateDtoToExchangeRateMapperImpl{})
	exchangeRateService.SetExchangeRateConvertDtoToExchangeRateResultMapper(mappers.ExchangeRateConvertDtoToExchangeRateConvertResultDtoMapperImpl{})
	exchangeRateService.SetExchangeRateApiService(services.ExchangeRateApiServiceImpl{})
	return exchangeRateService
}

func ScheduleUpdateRatesTask(exchangeRateService services_interfaces.ExchangeRateService, minutesCount int) {
	taskService := services.TaskServiceImpl{}
	taskService.SetExchangeRateService(exchangeRateService)
	taskService.RunUpdateExchangeRatesTask(minutesCount)
}

func RunHttpServer(exchangeRateService services_interfaces.ExchangeRateService) {
	httpServer := martini.Classic()
	httpServer.Use(render.Renderer(render.Options{
		Charset: "UTF-8",
	}))

	httpServer.Post("/api/create",
		binding.Json(dto.ExchangeRateAddInDatabaseDto{}),
		binding.ErrorHandler,
		func(params martini.Params, exchangeRate dto.ExchangeRateAddInDatabaseDto, jsonRenderer render.Render) {
			jsonRenderer.JSON(200, exchangeRateService.Add(exchangeRate))
		})

	httpServer.Post("/api/convert",
		binding.Json(dto.ExchangeRateConvertDto{}),
		binding.ErrorHandler,
		func(params martini.Params, exchangeRate dto.ExchangeRateConvertDto, jsonRenderer render.Render) {
			jsonRenderer.JSON(200, exchangeRateService.Convert(exchangeRate))
		})

	httpServer.Run()
}
