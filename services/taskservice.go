package services

type TaskService interface {
	RunUpdateExchangeRatesTask(minutesCount int)
}
