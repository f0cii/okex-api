package okex

import (
	"github.com/spf13/viper"
	"log"
	"testing"
)

func newFuturesWSForTest() *FuturesWS {
	viper.SetConfigName("test_config")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	accessKey := viper.GetString("access_key")
	secretKey := viper.GetString("secret_key")
	passphrase := viper.GetString("passphrase")
	wsURL := "wss://real.okex.com:8443/ws/v3"
	ws := NewFuturesWS(wsURL,
		accessKey, secretKey, passphrase)
	err = ws.SetProxy("socks5://127.0.0.1:1080")
	//ws.SetProxy("http://127.0.0.1:1080")
	if err != nil {
		log.Fatal(err)
	}
	return ws
}

func TestFuturesWS_AllInOne(t *testing.T) {
	ws := newFuturesWSForTest()
	//ws.SubscribeTicker("ticker_1", "BTC-USD-200626")
	//ws.SubscribeTrade("trade_1", "BTC-USD-200626")
	//ws.SubscribeDepthL2Tbt("depthL2_1", "BTC-USD-200626")
	//ws.SubscribeOrder("order_1", "BTC-USD-200626")
	ws.SubscribeOrder("position_1", "BTC-USD-200626")
	ws.Start()

	select {}
}
