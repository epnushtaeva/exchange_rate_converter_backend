package impl

import (
	"github.com/epnushtaeva/exchange_rate_converter_backend/services"
	"fmt"
	"time"
)

type TaskServiceImpl struct {
	exchangeRateService services.ExchangeRateService
}

func (taskServiceImpl *TaskServiceImpl) SetExchangeRateService(exchangeRateService services.ExchangeRateService) {
	taskServiceImpl.exchangeRateService = exchangeRateService
}

func (taskServiceImpl *TaskServiceImpl) RunUpdateExchangeRatesTask(minutesCount int) {
	tick := time.NewTicker(time.Minute * time.Duration(minutesCount))
	go taskServiceImpl.schedule(tick)
}

func (taskServiceImpl *TaskServiceImpl) schedule(tick *time.Ticker) {
	for t := range tick.C {
		fmt.Println("Running update rates")
		taskServiceImpl.exchangeRateService.UpdateAll(t)
	}
}
