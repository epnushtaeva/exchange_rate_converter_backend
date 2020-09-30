package dto

type ExchangeRateConvertResultDto struct {
	CurrencyFrom string  `json:"currencyFrom" binding:"required"`
	CurrencyTo   string  `json:"currencyTo" binding:"required"`
	Value        float32 `json:"value" binding:"required"`
	Course       float32 `json:"course" binding:"required"`
	Result       float32 `json:"result" binding:"required"`
}

func (exchangeRateConvertResultDto *ExchangeRateConvertResultDto) GetCurrencyFrom() string {
	return exchangeRateConvertResultDto.CurrencyFrom
}

func (exchangeRateConvertResultDto *ExchangeRateConvertResultDto) SetCurrencyFrom(currencyFrom string) {
	exchangeRateConvertResultDto.CurrencyFrom = currencyFrom
}

func (exchangeRateConvertResultDto *ExchangeRateConvertResultDto) GetCurrencyTo() string {
	return exchangeRateConvertResultDto.CurrencyTo
}

func (exchangeRateConvertResultDto *ExchangeRateConvertResultDto) SetCurrencyTo(currencyTo string) {
	exchangeRateConvertResultDto.CurrencyTo = currencyTo
}

func (exchangeRateConvertResultDto *ExchangeRateConvertResultDto) GetValue() float32 {
	return exchangeRateConvertResultDto.Value
}

func (exchangeRateConvertResultDto *ExchangeRateConvertResultDto) SetValue(value float32) {
	exchangeRateConvertResultDto.Value = value
}

func (exchangeRateConvertResultDto *ExchangeRateConvertResultDto) GetCourse() float32 {
	return exchangeRateConvertResultDto.Course
}

func (exchangeRateConvertResultDto *ExchangeRateConvertResultDto) SetCourse(course float32) {
	exchangeRateConvertResultDto.Course = course
}

func (exchangeRateConvertResultDto *ExchangeRateConvertResultDto) GetResult() float32 {
	return exchangeRateConvertResultDto.Result
}

func (exchangeRateConvertResultDto *ExchangeRateConvertResultDto) SetResult(result float32) {
	exchangeRateConvertResultDto.Result = result
}
