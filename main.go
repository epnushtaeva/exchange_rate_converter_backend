package main

import (
	"github.com/epnushtaeva/exchange_rate_converter_backend/internal"
)

func main() {
	dataBaseFactory := internal.InitDataBaseFactory("host=localhost user=postgres dbname=exchange_rate_database password=android sslmode=disable")
	exchangeRateService := internal.InitExchangeRateService(dataBaseFactory)

	internal.ScheduleUpdateRatesTask(exchangeRateService, 2)
	internal.RunHttpServer(exchangeRateService)
}
