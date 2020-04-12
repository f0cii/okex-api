package okex

import "testing"

func TestFuturesWS_AllInOne(t *testing.T) {
	wsURL := "wss://real.okex.com:8443/ws/v3"
	ws := NewFuturesWS(wsURL,
		"", "", "")
	err := ws.SetProxy("socks5://127.0.0.1:1080")
	//err := ws.SetProxy("http://127.0.0.1:10809")
	if err != nil {
		t.Error(err)
		return
	}
	//ws.SetProxy("http://127.0.0.1:1080")
	//ws.SubscribeTicker("ticker_1", "BTC-USD-200626")
	//ws.SubscribeTrade("trade_1", "BTC-USD-200626")
	//ws.SubscribeDepthL2Tbt("depthL2_1", "BTC-USD-200626")
	ws.SubscribeOrder("order_1", "BTC-USD-200626")
	ws.Start()

	select {}
}
