package okex

import (
	"fmt"
	"time"
)

type FinancialRate struct {
	ID     int    `json:"id"`
	Rate   string `json:"rate"`
	Symbol string `json:"symbol"`
}

type FinancialRates struct {
	Date  int64           `json:"date"`
	Rates []FinancialRate `json:"rates"`
}

type FinancialRatesResult struct {
	Code      int            `json:"code"`
	Data      FinancialRates `json:"data"`
	DetailMsg string         `json:"detailMsg"`
	Msg       string         `json:"msg"`
}

// https://www.okex.me/v2/asset/outer/financial/rates?t=1581044585771
func (client *Client) GetFinancialRates() (FinancialRatesResult, error) {
	// https://www.okex.me
	uri := fmt.Sprintf("/v2/asset/outer/financial/rates?t=%v", // 1581044585771
		time.Now().UnixNano()/1000000,
	)
	var result FinancialRatesResult
	_, _, err := client.Request(GET, uri, nil, &result)
	return result, err
}
