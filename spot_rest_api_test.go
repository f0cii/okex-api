package okex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSpotAccounts(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotAccounts()

	fmt.Printf("%+v, %+v", ac, err)

	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetSpotAccountsCurrency(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotAccountsCurrency("BTC")
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetSpotAccountsCurrencyLeger(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotAccountsCurrencyLeger("btc", nil)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)

	options := map[string]string{}
	options["from"] = "1"
	options["to"] = "2"
	options["limit"] = "100"

	ac2, err2 := c.GetSpotAccountsCurrencyLeger("btc", &options)
	assert.True(t, ac2 != nil && err2 == nil)
}

func TestGetSpotOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrders("filled", "BTC-USDT", nil)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)

	// Fore. 20190305. TODO: {"message":"System error"} returned by following request.
	// Url: http://coinmainweb.new.docker.okex.com/api/spot/v3/fills?instrument_id=BTC-USDT&order_id=2365709152770048
	filledOrderId := (*ac)[0]["order_id"].(string)
	sf, err := c.GetSpotFills(filledOrderId, "BTC-USDT", nil)
	assert.True(t, sf != nil && err == nil)
}

func TestGetSpotOrdersPending(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrdersPending(nil)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)

	options := NewParams()
	options["instrument_id"] = "BTC-USDT"
	ac, err = c.GetSpotOrdersPending(&options)
	assert.True(t, err == nil)
	jstr, _ = Struct2JsonString(ac)
	println(jstr)

	testOrderId := (*ac)[0]["order_id"]
	_, err = c.GetSpotOrdersById("BTC-USDT", testOrderId.(string))
	assert.True(t, err == nil)
}

func TestGetSpotInstruments(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotInstruments()
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetSpotInstrumentBook(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotInstrumentBook("LTC-USDT", nil)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetSpotInstrumentsTicker(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotInstrumentsTicker()
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetSpotInstrumentTicker(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotInstrumentTicker("LTC-USDT")
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetSpotInstrumentTrade(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotInstrumentTrade("BTC-USDT", nil)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)

	options := map[string]string{}
	options["from"] = "1"
	options["to"] = "2"
	options["limit"] = "100"

	ac2, err := c.GetSpotInstrumentTrade("BTC-USDT", &options)
	assert.True(t, err == nil)
	jstr, _ = Struct2JsonString(ac2)
	println(jstr)
}

func TestGetSpotInstrumentCandles(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotInstrumentCandles("BTC-USDT", nil)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestPostSpotOrders(t *testing.T) {
	c := NewTestClient()

	optionals := NewParams()
	optionals["type"] = "limit"
	optionals["price"] = "100"
	optionals["size"] = "0.01"

	r0, err := c.PostSpotOrders("sell", "btc-usdt", &optionals)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(r0)
	println(jstr)

	orderId := r0.OrderID
	r, err := c.PostSpotCancelOrders("btc-usdt", orderId)
	assert.True(t, r != nil && err == nil)
	jstr, _ = Struct2JsonString(r)
	println(jstr)

}

func TestClient_PostSpotBatchOrders(t *testing.T) {
	c := NewTestClient()

	orderInfos := []map[string]string{
		map[string]string{"client_oid": "w20180728w", "instrument_id": "btc-usdt", "side": "sell", "type": "limit", "size": "0.001", "price": "10001", "margin_trading ": "1"},
		map[string]string{"client_oid": "r20180728r", "instrument_id": "btc-usdt", "side": "sell", "type": "limit", " size ": "0.001", "notional": "10002", "margin_trading ": "1"},
	}

	r, err := c.PostSpotBatchOrders(&orderInfos)
	assert.True(t, r != nil && err == nil)
	jstr, _ := Struct2JsonString(r)
	println(jstr)
}

func TestClient_PostSpotCancelBatchOrders(t *testing.T) {
	c := NewTestClient()

	orderInfos := []map[string]interface{}{
		map[string]interface{}{"instrument_id": "btc-usdt", "client_oid": "16ee593327162368"},
		map[string]interface{}{"instrument_id": "ltc-usdt", "client_oid": "243464oo234465"},
	}

	r, err := c.PostSpotCancelBatchOrders(&orderInfos)
	assert.True(t, r != nil && err == nil)
	jstr, _ := Struct2JsonString(r)
	println(jstr)
}
