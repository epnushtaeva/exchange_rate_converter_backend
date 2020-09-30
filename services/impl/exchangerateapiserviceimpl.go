package impl

import (
	"encoding/json"
	"github.com/epnushtaeva/exchange_rate_converter_backend/dto"
	"io/ioutil"
	"net/http"
	"strings"
)

type ExchangeRateApiServiceImpl struct {
}

func (exchangeRateApiServiceImpl ExchangeRateApiServiceImpl) GetCourse(exchangeRateDto dto.ExchangeRateAddInDatabaseDto) float32 {
	exchangeRateName := strings.ToTitle(exchangeRateDto.GetCurrency1()) + "_" + strings.ToTitle(exchangeRateDto.GetCurrency2())
	exchangeRate, _ := http.Get("https://free.currconv.com/api/v7/convert?q=" +
		exchangeRateName +
		"&compact=ultra&apiKey=7d9b4f8a672d2285f2fd")
	defer exchangeRate.Body.Close()

	exchangeRateBytes, _ := ioutil.ReadAll(exchangeRate.Body)
	var exchangeRateMap map[string]float32
	json.Unmarshal(exchangeRateBytes, &exchangeRateMap)

	if course, ok := exchangeRateMap[exchangeRateName]; ok {
		return course
	}

	return 0
}
