package entities

import (
	"time"
)

type ExchangeRate struct {
	Id           int `gorm:"primaryKey;autoIncrement:true"`
	CurrencyFrom string
	CurrencyTo   string
	Course       float32
	UpdateDate   time.Time
}

func (exchangeRate *ExchangeRate) GetId() int {
	return exchangeRate.Id
}

func (exchangeRate *ExchangeRate) SetId(id int) *ExchangeRate {
	exchangeRate.Id = id
	return exchangeRate
}

func (exchangeRate *ExchangeRate) GetCurrencyFrom() string {
	return exchangeRate.CurrencyFrom
}

func (exchangeRate *ExchangeRate) SetCurrencyFrom(currencyFrom string) *ExchangeRate {
	exchangeRate.CurrencyFrom = currencyFrom
	return exchangeRate
}

func (exchangeRate *ExchangeRate) GetCurrencyTo() string {
	return exchangeRate.CurrencyTo
}

func (exchangeRate *ExchangeRate) SetCurrencyTo(currencyTo string) *ExchangeRate {
	exchangeRate.CurrencyTo = currencyTo
	return exchangeRate
}

func (exchangeRate *ExchangeRate) GetCourse() float32 {
	return exchangeRate.Course
}

func (exchangeRate *ExchangeRate) SetCourse(course float32) *ExchangeRate {
	exchangeRate.Course = course
	return exchangeRate
}

func (exchangeRate *ExchangeRate) GetUpdateDate() time.Time {
	return exchangeRate.UpdateDate
}

func (exchangeRate *ExchangeRate) SetUpdateDate(updateDate time.Time) *ExchangeRate {
	exchangeRate.UpdateDate = updateDate
	return exchangeRate
}
