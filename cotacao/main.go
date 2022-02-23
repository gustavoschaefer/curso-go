package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type resStock struct {
	Stocks []struct {
		Change       float64 `json:"change"`
		ClosingPrice float64 `json:"closingPrice"`
		Eps          float64 `json:"eps"`
		High         float64 `json:"high"`
		LastPrice    float64 `json:"lastPrice"`
		LastYearHigh float64 `json:"lastYearHigh"`
		LastYearLow  float64 `json:"lastYearLow"`
		Low          float64 `json:"low"`
		MarketCap    int     `json:"marketCap"`
		Name         string  `json:"name"`
		Pe           float64 `json:"pe"`
		PriceOpen    float64 `json:"priceOpen"`
		Shares       int     `json:"shares"`
		Symbol       string  `json:"symbol"`
		Volume       int     `json:"volume"`
		VolumeAvg    int     `json:"volumeAvg"`
		Sector       string  `json:"sector"`
		SubSector    string  `json:"subSector"`
		Segment      string  `json:"segment"`
	} `json:"stocks"`
}

type resStockUnit struct {
	Change       int     `json:"change"`
	ClosingPrice float64 `json:"closingPrice"`
	Eps          float64 `json:"eps"`
	High         float64 `json:"high"`
	LastPrice    float64 `json:"lastPrice"`
	LastYearHigh float64 `json:"lastYearHigh"`
	LastYearLow  float64 `json:"lastYearLow"`
	Low          float64 `json:"low"`
	MarketCap    int64   `json:"marketCap"`
	Name         string  `json:"name"`
	Pe           float64 `json:"pe"`
	PriceOpen    float64 `json:"priceOpen"`
	Shares       int64   `json:"shares"`
	Symbol       string  `json:"symbol"`
	Volume       int     `json:"volume"`
	VolumeAvg    int     `json:"volumeAvg"`
	Sector       string  `json:"sector"`
	SubSector    string  `json:"subSector"`
	Segment      string  `json:"segment"`
}

type resFiis struct {
	Change       int     `json:"change"`
	ClosingPrice float64 `json:"closingPrice"`
	High         float64 `json:"high"`
	LastPrice    float64 `json:"lastPrice"`
	LastYearHigh float64 `json:"lastYearHigh"`
	LastYearLow  float64 `json:"lastYearLow"`
	Low          float64 `json:"low"`
	Name         string  `json:"name"`
	PriceOpen    float64 `json:"priceOpen"`
	Shares       int     `json:"shares"`
	Symbol       string  `json:"symbol"`
	Volume       int     `json:"volume"`
	VolumeAvg    int     `json:"volumeAvg"`
}

func main() {
	fmt.Println("Buscando ativos...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://mfinance.com.br/api/v1/stocks/b3sa3", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var respObj resStock
	json.Unmarshal(bodyBytes, &respObj)

	if respObj.Stocks == nil {
		var respObjUnit resStockUnit
		json.Unmarshal(bodyBytes, &respObjUnit)
		fmt.Println(respObjUnit.Symbol, respObjUnit.LastPrice)
	} else {
		for _, value := range respObj.Stocks {
			fmt.Println(value.Symbol, value.LastPrice)
		}
	}

}
