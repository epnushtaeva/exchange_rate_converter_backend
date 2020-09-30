package dto

type ExchangeRateAddInDatabaseDto struct {
	Currency1 string `json:"currency1" binding:"required"`
	Currency2 string `json:"currency2" binding:"required"`
}

func (exchangeRateAddInDataBaseDto *ExchangeRateAddInDatabaseDto) GetCurrency1() string {
	return exchangeRateAddInDataBaseDto.Currency1
}

func (exchangeRateAddInDataBaseDto *ExchangeRateAddInDatabaseDto) SetCurrency1(currency1 string) {
	exchangeRateAddInDataBaseDto.Currency1 = currency1
}

func (exchangeRateAddInDataBaseDto *ExchangeRateAddInDatabaseDto) GetCurrency2() string {
	return exchangeRateAddInDataBaseDto.Currency2
}

func (exchangeRateAddInDataBaseDto *ExchangeRateAddInDatabaseDto) SetCurrency2(currency2 string) {
	exchangeRateAddInDataBaseDto.Currency2 = currency2
}
