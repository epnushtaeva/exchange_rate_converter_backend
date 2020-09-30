package dto

type ExchangeRateConvertDto struct {
	CurrencyFrom string  `json:"currencyFrom" binding:"required"`
	CurrencyTo   string  `json:"currencyTo" binding:"required"`
	Value        float32 `json:"value" binding:"required"`
}

func (exchangeRateConvertDto *ExchangeRateConvertDto) GetCurrencyFrom() string {
	return exchangeRateConvertDto.CurrencyFrom
}

func (exchangeRateConvertDto *ExchangeRateConvertDto) SetCurrencyFrom(currencyFrom string) {
	exchangeRateConvertDto.CurrencyFrom = currencyFrom
}

func (exchangeRateConvertDto *ExchangeRateConvertDto) GetCurrencyTo() string {
	return exchangeRateConvertDto.CurrencyTo
}

func (exchangeRateConvertDto *ExchangeRateConvertDto) SetCurrencyTo(currencyTo string) {
	exchangeRateConvertDto.CurrencyTo = currencyTo
}

func (exchangeRateConvertDto *ExchangeRateConvertDto) GetValue() float32 {
	return exchangeRateConvertDto.Value
}

func (exchangeRateConvertDto *ExchangeRateConvertDto) SetValue(value float32) {
	exchangeRateConvertDto.Value = value
}
