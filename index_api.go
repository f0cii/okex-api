package okex

import (
	"time"
)

type Constituent struct {
	Symbol        string  `json:"symbol"`
	OriginalPrice float64 `json:"original_price,string"`
	Weight        float64 `json:"weight,string"`
	UsdPrice      float64 `json:"usd_price,string"`
	Exchange      string  `json:"exchange"`
}

type IndexConstituents struct {
	Last         string        `json:"last"`
	Constituents []Constituent `json:"constituents"`
	InstrumentID string        `json:"instrument_id"`
	Timestamp    time.Time     `json:"timestamp"`
}

type IndexConstituentsResult struct {
	Code      int               `json:"code"`
	Data      IndexConstituents `json:"data"`
	DetailMsg string            `json:"detailMsg"`
	Msg       string            `json:"msg"`
}

// GetIndexConstituents 公共-获取指数成分
func (client *Client) GetIndexConstituents(instrumentID string) (IndexConstituentsResult, error) {
	baseURI := GetInstrumentIdUri(INDEX_CONSTITUENTS, instrumentID)
	var result IndexConstituentsResult
	_, _, err := client.Request(GET, baseURI, nil, &result)
	// if err != nil {
	// 	return result, err
	// }
	// body, _ := ioutil.ReadAll(resp.Body)
	// log.Printf("%v", string(body))
	return result, err
}
