package okex

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/recws-org/recws"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	TableSwapTicker     = "swap/ticker"       // 公共-Ticker频道
	TableSwapTrade      = "swap/trade"        // 公共-交易频道
	TableSwapDepthL2Tbt = "swap/depth_l2_tbt" // 公共-400档增量数据频道
	TableSwapPosition   = "swap/position"     // 用户持仓频道
	TableSwapAccount    = "swap/account"      // 用户账户频道
	TableSwapOrder      = "swap/order"        // 用户交易频道
)

type SwapWS struct {
	sync.RWMutex

	wsURL      string
	accessKey  string
	secretKey  string
	passphrase string

	ctx    context.Context
	cancel context.CancelFunc
	wsConn recws.RecConn

	subscriptions map[string]interface{}

	tickersCallback    func(tickers []WSTicker)
	tradesCallback     func(trades []WSTrade)
	depthL2TbtCallback func(action string, data []WSDepthL2Tbt)
	accountCallback    func(accounts []WSAccount)
	positionCallback   func(positions []WSSwapPositionData)
	orderCallback      func(orders []WSOrder)
}

// SetProxy 设置代理地址
// porxyURL:
// socks5://127.0.0.1:1080
// https://127.0.0.1:1080
func (ws *SwapWS) SetProxy(proxyURL string) (err error) {
	var purl *url.URL
	purl, err = url.Parse(proxyURL)
	if err != nil {
		return
	}
	log.Printf("[ws][%s] proxy url:%s", proxyURL, purl)
	ws.wsConn.Proxy = http.ProxyURL(purl)
	return
}

func (ws *SwapWS) SetTickerCallback(callback func(tickers []WSTicker)) {
	ws.tickersCallback = callback
}

func (ws *SwapWS) SetTradeCallback(callback func(trades []WSTrade)) {
	ws.tradesCallback = callback
}

func (ws *SwapWS) SetDepthL2TbtCallback(callback func(action string, data []WSDepthL2Tbt)) {
	ws.depthL2TbtCallback = callback
}

func (ws *SwapWS) SetAccountCallback(callback func(accounts []WSAccount)) {
	ws.accountCallback = callback
}

func (ws *SwapWS) SetPositionCallback(callback func(position []WSSwapPositionData)) {
	ws.positionCallback = callback
}

func (ws *SwapWS) SetOrderCallback(callback func(orders []WSOrder)) {
	ws.orderCallback = callback
}

func (ws *SwapWS) SubscribeTicker(id string, symbol string) error {
	ch := fmt.Sprintf("%v:%v", TableSwapTicker, symbol)
	return ws.Subscribe(id, []string{ch})
}

func (ws *SwapWS) SubscribeTrade(id string, symbol string) error {
	ch := fmt.Sprintf("%v:%v", TableSwapTrade, symbol)
	return ws.Subscribe(id, []string{ch})
}

// SubscribeDepthL2Tbt 公共-400档增量数据频道
// 订阅后首次返回市场订单簿的400档深度数据并推送；后续只要订单簿深度有变化就推送有更改的数据。
func (ws *SwapWS) SubscribeDepthL2Tbt(id string, symbol string) error {
	ch := fmt.Sprintf("%v:%v", TableSwapDepthL2Tbt, symbol)
	return ws.Subscribe(id, []string{ch})
}

func (ws *SwapWS) SubscribePosition(id string, symbol string) error {
	ch := fmt.Sprintf("%v:%v", TableSwapPosition, symbol)
	return ws.Subscribe(id, []string{ch})
}

func (ws *SwapWS) SubscribeAccount(id string, symbol string) error {
	ch := fmt.Sprintf("%v:%v", TableSwapAccount, symbol)
	return ws.Subscribe(id, []string{ch})
}

func (ws *SwapWS) SubscribeOrder(id string, symbol string) error {
	ch := fmt.Sprintf("%v:%v", TableSwapOrder, symbol)
	return ws.Subscribe(id, []string{ch})
}

// Subscribe 订阅
func (ws *SwapWS) Subscribe(id string, args []string) error {
	ws.Lock()
	defer ws.Unlock()

	type Op struct {
		Op   string   `json:"op"`
		Args []string `json:"args"`
	}

	op := Op{
		Op:   "subscribe",
		Args: args,
	}
	ws.subscriptions[id] = op
	return ws.sendWSMessage(op)
}

// Unsubscribe 取消订阅
func (ws *SwapWS) Unsubscribe(id string) error {
	ws.Lock()
	defer ws.Unlock()

	if _, ok := ws.subscriptions[id]; ok {
		delete(ws.subscriptions, id)
	}
	return nil
}

func (ws *SwapWS) Login() error {
	if ws.accessKey == "" || ws.secretKey == "" || ws.passphrase == "" {
		return fmt.Errorf("missing key")
	}
	timestamp := EpochTime()

	preHash := PreHashString(timestamp, GET, "/users/self/verify", "")
	if sign, err := HmacSha256Base64Signer(preHash, ws.secretKey); err != nil {
		return err
	} else {
		op, err := loginOp(ws.accessKey, ws.passphrase, timestamp, sign)
		if err != nil {
			return err
		}
		//data, err := Struct2JsonString(op)
		log.Printf("Send Msg: %#v", *op)
		//err = a.conn.WriteMessage(websocket.TextMessage, []byte(data))
		err = ws.sendWSMessage(op)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 100)
	}
	return nil
}

func (ws *SwapWS) subscribeHandler() error {
	//log.Printf("subscribeHandler")
	ws.Lock()
	defer ws.Unlock()

	err := ws.Login()
	if err != nil {
		log.Printf("login error: %v", err)
	}

	for _, v := range ws.subscriptions {
		//log.Printf("sub: %#v", v)
		err := ws.sendWSMessage(v)
		if err != nil {
			log.Printf("%v", err)
		}
	}
	return nil
}

func (ws *SwapWS) sendWSMessage(msg interface{}) error {
	return ws.wsConn.WriteJSON(msg)
}

func (ws *SwapWS) Start() {
	log.Printf("wsURL: %v", ws.wsURL)
	ws.wsConn.Dial(ws.wsURL, nil)
	go ws.run()
}

func (ws *SwapWS) run() {
	ctx := context.Background()
	for {
		select {
		case <-ctx.Done():
			go ws.wsConn.Close()
			log.Printf("Websocket closed %s", ws.wsConn.GetURL())
			return
		default:
			messageType, msg, err := ws.wsConn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				time.Sleep(100 * time.Millisecond)
				continue
			}

			msg, err = FlateUnCompress(msg)
			if err != nil {
				log.Printf("%v", err)
				continue
			}

			ws.handleMsg(messageType, msg)
		}
	}
}

func (ws *SwapWS) handleMsg(messageType int, msg []byte) {
	ret := gjson.ParseBytes(msg)
	// 登录成功
	// {"event":"login","success":true}

	if tableValue := ret.Get("table"); tableValue.Exists() {
		table := tableValue.String()
		if table == TableSwapDepthL2Tbt { // 优先判断最高频数据
			var depthL2 WSDepthL2TbtResult
			err := json.Unmarshal(msg, &depthL2)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			if ws.depthL2TbtCallback != nil {
				ws.depthL2TbtCallback(depthL2.Action, depthL2.Data)
			}
			return
		} else if table == TableSwapTicker {
			var tickerResult WSTickerResult
			err := json.Unmarshal(msg, &tickerResult)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			if ws.tickersCallback != nil {
				ws.tickersCallback(tickerResult.Data)
			}
			return
		} else if table == TableSwapTrade {
			var tradeResult WSTradeResult
			err := json.Unmarshal(msg, &tradeResult)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			if ws.tradesCallback != nil {
				ws.tradesCallback(tradeResult.Data)
			}
			return
		} else if table == TableSwapAccount {
			var accountResult WSAccountResult
			err := json.Unmarshal(msg, &accountResult)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			if ws.accountCallback != nil {
				var accounts []WSAccount
				for _, v := range accountResult.Data {
					if v.BTC != nil {
						accounts = append(accounts, *v.BTC)
						continue
					}
					if v.ETH != nil {
						accounts = append(accounts, *v.ETH)
						continue
					}
					if v.ETC != nil {
						accounts = append(accounts, *v.ETC)
						continue
					}
					if v.XRP != nil {
						accounts = append(accounts, *v.XRP)
						continue
					}
					if v.EOS != nil {
						accounts = append(accounts, *v.EOS)
						continue
					}
					if v.BCH != nil {
						accounts = append(accounts, *v.BCH)
						continue
					}
					if v.BSV != nil {
						accounts = append(accounts, *v.BSV)
						continue
					}
					if v.TRX != nil {
						accounts = append(accounts, *v.TRX)
						continue
					}
				}
				ws.accountCallback(accounts)
			}
			return
		} else if table == TableSwapPosition {
			var positionResult WSSwapPositionResult
			err := json.Unmarshal(msg, &positionResult)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			if ws.positionCallback != nil {
				ws.positionCallback(positionResult.Data)
			}
			return
		} else if table == TableSwapOrder {
			var orderResult WSOrderResult
			err := json.Unmarshal(msg, &orderResult)
			if err != nil {
				log.Printf("%v", err)
				return
			}

			if ws.orderCallback != nil {
				ws.orderCallback(orderResult.Data)
			}
			return
		}
		log.Printf("%v", string(msg))
		return
	}

	if eventValue := ret.Get("event"); eventValue.Exists() {
		event := eventValue.String()
		if event == "error" {
			log.Printf("error: %v", string(msg))
			return
		}
		log.Printf("%v", string(msg))
		return
	}

	log.Printf("%v", string(msg))
}

// NewSwapWS 创建永续合约WS
// wsURL:
// wss://real.okex.com:8443/ws/v3
func NewSwapWS(wsURL string, accessKey string, secretKey string, passphrase string) *SwapWS {
	ws := &SwapWS{
		wsURL:         wsURL,
		accessKey:     accessKey,
		secretKey:     secretKey,
		passphrase:    passphrase,
		subscriptions: make(map[string]interface{}),
	}
	ws.ctx, ws.cancel = context.WithCancel(context.Background())
	ws.wsConn = recws.RecConn{
		KeepAliveTimeout: 10 * time.Second,
	}
	ws.wsConn.SubscribeHandler = ws.subscribeHandler
	return ws
}
