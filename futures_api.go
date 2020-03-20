package okex

import (
	"log"
	"net/http"
	"strings"
)

/*
 OKEX futures contract api
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

/*
 =============================== Futures market api ===============================
*/
/*
 The exchange rate of legal tender pairs
*/
func (client *Client) GetFuturesExchangeRate() (ExchangeRate, error) {
	var exchangeRate ExchangeRate
	_, _, err := client.Request(GET, FUTURES_RATE, nil, &exchangeRate)
	return exchangeRate, err
}

/*
  Get all of futures contract list
*/
func (client *Client) GetFuturesInstruments() ([]FuturesInstrumentsResult, error) {
	var Instruments []FuturesInstrumentsResult
	_, _, err := client.Request(GET, FUTURES_INSTRUMENTS, nil, &Instruments)
	return Instruments, err
}

/*
 Get the futures contract currencies
*/
func (client *Client) GetFuturesInstrumentCurrencies() ([]FuturesInstrumentCurrenciesResult, error) {
	var currencies []FuturesInstrumentCurrenciesResult
	_, _, err := client.Request(GET, FUTURES_CURRENCIES, nil, &currencies)
	return currencies, err
}

/*
	获取深度数据
	获取币对的深度列表。这个请求不支持分页，一个请求返回整个深度列表。

	限速规则：20次/2s
	HTTP请求
	GET /api/spot/v3/instruments/<instrument_id>/book

	签名请求示例
	2018-09-12T07:57:09.130ZGET/api/spot/v3/instruments/LTC-USDT/book?size=10&depth=0.001

*/
func (client *Client) GetFuturesInstrumentBook(InstrumentId string, optionalParams map[string]string) (FuturesInstrumentBookResult, error) {
	var book FuturesInstrumentBookResult
	params := NewParams()
	if optionalParams != nil && len(optionalParams) > 0 {
		params["size"] = optionalParams["size"]
		params["depth"] = optionalParams["depth"]
	}
	requestPath := BuildParams(GetInstrumentIdUri(FUTURES_INSTRUMENT_BOOK, InstrumentId), params)
	_, _, err := client.Request(GET, requestPath, nil, &book)
	return book, err
}

/*
 Get the futures contract Instrument all ticker
*/
func (client *Client) GetFuturesInstrumentAllTicker() ([]FuturesInstrumentTickerResult, error) {
	var tickers []FuturesInstrumentTickerResult
	_, _, err := client.Request(GET, FUTURES_TICKERS, nil, &tickers)
	return tickers, err
}

/*
 Get the futures contract Instrument ticker
*/
func (client *Client) GetFuturesInstrumentTicker(InstrumentId string) (FuturesInstrumentTickerResult, error) {
	var ticker FuturesInstrumentTickerResult
	_, _, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_TICKER, InstrumentId), nil, &ticker)
	return ticker, err
}

/*
 Get the futures contract Instrument trades
*/
func (client *Client) GetFuturesInstrumentTrades(InstrumentId string) ([]FuturesInstrumentTradesResult, error) {
	var trades []FuturesInstrumentTradesResult
	_, _, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_TRADES, InstrumentId), nil, &trades)
	return trades, err
}

/*
 Get the futures contract Instrument candles
 granularity: @see  file: futures_constants.go
*/
func (client *Client) GetFuturesInstrumentCandles(InstrumentId string, optionalParams map[string]string) ([][]string, error) {
	var candles [][]string
	params := NewParams()

	if optionalParams != nil && len(optionalParams) > 0 {
		params["start"] = optionalParams["start"]
		params["end"] = optionalParams["end"]
		params["granularity"] = optionalParams["granularity"]
	}
	requestPath := BuildParams(GetInstrumentIdUri(FUTURES_INSTRUMENT_CANDLES, InstrumentId), params)
	_, _, err := client.Request(GET, requestPath, nil, &candles)
	return candles, err
}

/*
 Get the futures contract Instrument index
*/
func (client *Client) GetFuturesInstrumentIndex(InstrumentId string) (FuturesInstrumentIndexResult, error) {
	var index FuturesInstrumentIndexResult
	_, _, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_INDEX, InstrumentId), nil, &index)
	return index, err
}

/*
 Get the futures contract Instrument estimated price
*/
func (client *Client) GetFuturesInstrumentEstimatedPrice(InstrumentId string) (FuturesInstrumentEstimatedPriceResult, error) {
	var estimatedPrice FuturesInstrumentEstimatedPriceResult
	_, _, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_ESTIMATED_PRICE, InstrumentId), nil, &estimatedPrice)
	return estimatedPrice, err
}

/*
 Get the futures contract Instrument holds
*/
func (client *Client) GetFuturesInstrumentOpenInterest(InstrumentId string) (FuturesInstrumentOpenInterestResult, error) {
	var openInterest FuturesInstrumentOpenInterestResult
	_, _, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_OPEN_INTEREST, InstrumentId), nil, &openInterest)
	return openInterest, err
}

/*
 Get the futures contract Instrument limit price
*/
func (client *Client) GetFuturesInstrumentPriceLimit(InstrumentId string) (FuturesInstrumentPriceLimitResult, error) {
	var priceLimit FuturesInstrumentPriceLimitResult
	_, _, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_PRICE_LIMIT, InstrumentId), nil, &priceLimit)
	return priceLimit, err
}

/*
 Get the futures contract liquidation
*/
func (client *Client) GetFuturesInstrumentLiquidation(InstrumentId string, status, from, to, limit int) (FuturesInstrumentLiquidationListResult, error) {
	var liquidation []FuturesInstrumentLiquidationResult
	params := NewParams()
	params["status"] = Int2String(status)
	params["from"] = Int2String(from)
	params["to"] = Int2String(to)
	params["limit"] = Int2String(limit)
	requestPath := BuildParams(GetInstrumentIdUri(FUTURES_INSTRUMENT_LIQUIDATION, InstrumentId), params)
	_, response, err := client.Request(GET, requestPath, nil, &liquidation)
	var list FuturesInstrumentLiquidationListResult
	page := parsePage(response)
	list.Page = page
	list.LiquidationList = liquidation
	return list, err
}

/*
 =============================== Futures trade api ===============================
*/

/*
 Get all of futures contract position list.
 return struct: FuturesPositions
*/
func (client *Client) GetFuturesPositions() (FuturesPosition, error) {
	_, response, err := client.Request(GET, FUTURES_POSITION, nil, nil)
	return parsePositions(response, err)
}

/*
 Get all of futures contract position list.
 return struct: FuturesPositions
*/
func (client *Client) GetFuturesInstrumentPosition(InstrumentId string) (FuturesPosition, error) {
	_, response, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_POSITION, InstrumentId), nil, nil)
	return parsePositions(response, err)
}

/*
 Get all of futures contract account list
 return struct: FuturesAccounts
*/
func (client *Client) GetFuturesAccounts() (GetFuturesAccountsResult, error) {
	var r GetFuturesAccountsResult
	_, _, err := client.Request(GET, FUTURES_ACCOUNTS, nil, &r)
	//return parseAccounts(response, err)
	return r, err
}

/*
 Get the futures contract currency account @see file : futures_constants.go
 return struct: FuturesCurrencyAccounts
*/
func (client *Client) GetFuturesAccountsByCurrency(currency string) (result FuturesCurrencyAccount, err error) {
	_, _, err = client.Request(GET, GetUnderlyingUri(FUTURES_ACCOUNT_CURRENCY_INFO, currency), nil, &result)
	//return parseCurrencyAccounts(response, err)
	return
}

/*
 Get the futures contract currency ledger
*/
func (client *Client) GetFuturesAccountsLedgerByCurrency(currency string, from, to, limit int) ([]FuturesCurrencyLedger, error) {
	var ledger []FuturesCurrencyLedger
	params := NewParams()
	params["from"] = Int2String(from)
	params["to"] = Int2String(to)
	params["limit"] = Int2String(limit)
	requestPath := BuildParams(GetCurrencyUri(FUTURES_ACCOUNT_CURRENCY_LEDGER, currency), params)
	_, _, err := client.Request(GET, requestPath, nil, &ledger)
	return ledger, err
}

/*
 Get the futures contract Instrument holds
*/
func (client *Client) GetFuturesAccountsHoldsByInstrumentId(InstrumentId string) (FuturesAccountsHolds, error) {
	var holds FuturesAccountsHolds
	_, _, err := client.Request(GET, GetInstrumentIdUri(FUTURES_ACCOUNT_INSTRUMENT_HOLDS, InstrumentId), nil, &holds)
	return holds, err
}

/*
 Create a new order
*/
func (client *Client) FuturesOrder(newOrderParams FuturesNewOrderParams) ([]byte, FuturesNewOrderResult, error) {
	var newOrderResult FuturesNewOrderResult
	var respBody []byte
	respBody, _, err := client.Request(POST, FUTURES_ORDER, newOrderParams, &newOrderResult)
	return respBody, newOrderResult, err
}

/*
 Batch create new order.(Max of 5 orders are allowed per request)
*/
func (client *Client) FuturesOrders(batchNewOrder FuturesBatchNewOrderParams) ([]byte, FuturesBatchNewOrderResult, error) {
	var batchNewOrderResult FuturesBatchNewOrderResult
	var respBody []byte
	respBody, _, err := client.Request(POST, FUTURES_ORDERS, batchNewOrder, &batchNewOrderResult)
	return respBody, batchNewOrderResult, err
}

/*
 Get all of futures contract order list
*/
func (client *Client) GetFuturesOrders(InstrumentId string, status, from, to, limit int) (FuturesGetOrdersResult, error) {
	var ordersResult FuturesGetOrdersResult
	params := NewParams()
	params["status"] = Int2String(status)
	params["from"] = Int2String(from)
	params["to"] = Int2String(to)
	params["limit"] = Int2String(limit)
	requestPath := BuildParams(GetInstrumentIdUri(FUTURES_INSTRUMENT_ORDER_LIST, InstrumentId), params)
	_, _, err := client.Request(GET, requestPath, nil, &ordersResult)
	return ordersResult, err
}

/*
 Get all of futures contract a order by order id
*/
func (client *Client) GetFuturesOrder(InstrumentId string, orderId string) (FuturesGetOrderResult, error) {
	var getOrderResult FuturesGetOrderResult
	_, _, err := client.Request(GET, GetInstrumentIdOrdersUri(FUTURES_INSTRUMENT_ORDER_INFO, InstrumentId, orderId), nil, &getOrderResult)
	return getOrderResult, err
}

/*
 Batch Cancel the orders
*/
func (client *Client) BatchCancelFuturesInstrumentOrders(InstrumentId, orderIds string) ([]byte, FuturesBatchCancelInstrumentOrdersResult, error) {
	var cancelInstrumentOrdersResult FuturesBatchCancelInstrumentOrdersResult
	params := NewParams()
	params["order_ids"] = orderIds
	var respBody []byte
	respBody, _, err := client.Request(POST, GetInstrumentIdUri(FUTURES_INSTRUMENT_ORDER_BATCH_CANCEL, InstrumentId), params, &cancelInstrumentOrdersResult)
	return respBody, cancelInstrumentOrdersResult, err
}

/*
 Cancel the order
*/
func (client *Client) CancelFuturesInstrumentOrder(InstrumentId string, orderId string) ([]byte, FuturesCancelInstrumentOrderResult, error) {
	var cancelInstrumentOrderResult FuturesCancelInstrumentOrderResult
	var respBody []byte
	respBody, _, err := client.Request(POST, GetInstrumentIdOrdersUri(FUTURES_INSTRUMENT_ORDER_CANCEL, InstrumentId, orderId), nil,
		&cancelInstrumentOrderResult)
	return respBody, cancelInstrumentOrderResult, err
}

/*
 Get all of futures contract transactions.
*/
func (client *Client) GetFuturesFills(InstrumentId string, orderId int64, optionalParams map[string]int) ([]FuturesFillResult, error) {
	var fillsResult []FuturesFillResult
	params := NewParams()
	params["order_id"] = Int64ToString(orderId)
	params["instrument_id"] = InstrumentId

	if optionalParams != nil && len(optionalParams) > 0 {
		params["from"] = Int2String(optionalParams["from"])
		params["to"] = Int2String(optionalParams["to"])
		params["limit"] = Int2String(optionalParams["limit"])
	}

	requestPath := BuildParams(FUTURES_FILLS, params)
	_, _, err := client.Request(GET, requestPath, nil, &fillsResult)
	return fillsResult, err
}

/*
获取标记价格
获取合约标记价格。此接口为公共接口，不需要身份验证。

请求示例
GET/api/futures/v3/instruments/BTC-USD-180309/mark_price
*/
func (client *Client) GetInstrumentMarkPrice(instrumentId string) (*FuturesMarkdown, error) {
	uri := GetInstrumentIdUri(FUTURES_INSTRUMENT_MARK_PRICE, instrumentId)
	r := FuturesMarkdown{}
	_, _, err := client.Request(GET, uri, nil, &r)
	return &r, err
}

func parsePositions(response *http.Response, err error) (FuturesPosition, error) {
	var position FuturesPosition
	if err != nil {
		return position, err
	}
	var result Result
	result.Result = false
	jsonString := GetResponseDataJsonString(response)
	if strings.Contains(jsonString, "\"margin_mode\":\"fixed\"") {
		var fixedPosition FuturesFixedPosition
		err = JsonString2Struct(jsonString, &fixedPosition)
		if err != nil {
			return position, err
		} else {
			position.Result = fixedPosition.Result
			position.MarginMode = fixedPosition.MarginMode
			position.FixedPosition = fixedPosition.FixedPosition
		}
	} else if strings.Contains(jsonString, "\"margin_mode\":\"crossed\"") {
		var crossPosition FuturesCrossPosition
		err = JsonString2Struct(jsonString, &crossPosition)
		if err != nil {
			return position, err
		} else {
			position.Result = crossPosition.Result
			position.MarginMode = crossPosition.MarginMode
			position.CrossPosition = crossPosition.CrossPosition
		}
	} else if strings.Contains(jsonString, "\"code\":") {
		JsonString2Struct(jsonString, &position)
		position.Result = result
	} else {
		position.Result = result
	}

	return position, nil
}

func parseAccounts(response *http.Response, err error) (FuturesAccount, error) {
	var account FuturesAccount
	if err != nil {
		return account, err
	}
	var result Result
	result.Result = false
	jsonString := GetResponseDataJsonString(response)
	log.Printf(jsonString)
	if strings.Contains(jsonString, "\"contracts\"") {
		var fixedAccount FuturesFixedAccountInfo
		err = JsonString2Struct(jsonString, &fixedAccount)
		if err != nil {
			return account, err
		} else {
			account.Result = fixedAccount.Result
			account.FixedAccount = fixedAccount.Info
			account.MarginMode = "fixed"
		}
	} else if strings.Contains(jsonString, "\"realized_pnl\"") {
		var crossAccount FuturesCrossAccountInfo
		err = JsonString2Struct(jsonString, &crossAccount)
		if err != nil {
			return account, err
		} else {
			account.Result = crossAccount.Result
			account.MarginMode = "crossed"
			account.CrossAccount = crossAccount.Info
		}
	} else if strings.Contains(jsonString, "\"code\":") {
		JsonString2Struct(jsonString, &account)
		account.Result = result
	} else {
		account.Result = result
	}
	return account, nil
}

func parseCurrencyAccounts(response *http.Response, err error) (FuturesCurrencyAccount, error) {
	var currencyAccount FuturesCurrencyAccount
	if err != nil {
		return currencyAccount, err
	}
	jsonString := GetResponseDataJsonString(response)
	var result Result
	result.Result = true
	if strings.Contains(jsonString, "\"margin_mode\":\"fixed\"") {
		var fixedAccount FuturesFixedAccount
		err = JsonString2Struct(jsonString, &fixedAccount)
		if err != nil {
			return currencyAccount, err
		} else {
			//currencyAccount.Result = result
			currencyAccount.MarginMode = fixedAccount.MarginMode
			//currencyAccount.FixedAccount = fixedAccount
		}
	} else if strings.Contains(jsonString, "\"margin_mode\":\"crossed\"") {
		var crossAccount FuturesCrossAccount
		err = JsonString2Struct(jsonString, &crossAccount)
		if err != nil {
			return currencyAccount, err
		} else {
			//currencyAccount.Result = result
			currencyAccount.MarginMode = crossAccount.MarginMode
			//currencyAccount.CrossAccount = crossAccount
		}
	} else if strings.Contains(jsonString, "\"code\":") {
		result.Result = true
		JsonString2Struct(jsonString, &currencyAccount)
		//currencyAccount.Result = result
	} else {
		result.Result = true
		//currencyAccount.Result = result
	}
	return currencyAccount, nil
}

func parsePage(response *http.Response) PageResult {
	var page PageResult
	jsonString := GetResponsePageJsonString(response)
	JsonString2Struct(jsonString, &page)
	return page
}

/*
设定合约币种杠杆倍数
设定合约账户币种杠杆倍数，注意当前仓位有持仓或者挂单禁止切换杠杆。

HTTP请求
POST /api/futures/v3/accounts/<currency>/leverage

请求示例
POST/api/futures/v3/accounts/btc/leverage{"leverage":"10"}（全仓示例）
POST/api/futures/v3/accounts/btc/leverage{"instrument_id":"BTC-USD-180213","direction":"long","leverage":"10"}（逐仓示例）

*/
func (client *Client) PostFuturesAccountsLeverage(currency string, leverage int, optionalParams map[string]string) (map[string]interface{}, error) {
	uri := GetUnderlyingUri(FUTURES_ACCOUNT_CURRENCY_LEVERAGE, currency)
	params := NewParams()
	params["leverage"] = Int2String(leverage)

	if optionalParams != nil && len(optionalParams) > 0 {
		params["instrument_id"] = optionalParams["instrument_id"]
		params["direction"] = optionalParams["direction"]
	}

	r := new(map[string]interface{})
	_, _, err := client.Request(POST, uri, params, r)

	return *r, err
}

/*
设置合约账户模式
请求示例
POST /api/futures/v3/accounts/margin_mode{"underlying":"btc-usd","margin_mode":"crossed"}
*/

func (client *Client) PostFuturesAccountsMarginNode(underlying string, marginMode string) (map[string]interface{}, error) {
	params := NewParams()
	params["underlying"] = underlying
	params["margin_mode"] = marginMode
	r := new(map[string]interface{})
	_, _, err := client.Request(POST, FUTURES_ACCOUNT_MARGIN_MODE, params, r)
	return *r, err
}

/*
获取合约账户币种杠杆倍数

限速规则：5次/2s
HTTP请求
GET /api/futures/v3/accounts/<currency>/leverage

请求示例
GET/api/futures/v3/accounts/btc/leverage
*/
func (client *Client) GetFuturesAccountsLeverage(currency string) (map[string]interface{}, error) {
	uri := GetUnderlyingUri(FUTURES_ACCOUNT_CURRENCY_LEVERAGE, currency)
	r := new(map[string]interface{})
	_, _, err := client.Request(GET, uri, nil, r)
	return *r, err
}
