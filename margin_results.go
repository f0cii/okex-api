package okex

import (
	sjson "encoding/json"
	"time"
)

type MarginCurrency struct {
	Available   string `json:"available"`
	Balance     string `json:"balance"`
	Borrowed    string `json:"borrowed"`
	CanWithdraw string `json:"can_withdraw"`
	Frozen      string `json:"frozen"`
	Hold        string `json:"hold"`
	Holds       string `json:"holds"`
	LendingFee  string `json:"lending_fee"`
}

type GetMarginAccountsItem struct {
	CurrencyBTC      MarginCurrency `json:"currency:BTC,omitempty"`
	CurrencyUSDT     MarginCurrency `json:"currency:USDT,omitempty"`
	InstrumentID     string         `json:"instrument_id"`
	LiquidationPrice string         `json:"liquidation_price"`
	MarginRatio      string         `json:"margin_ratio"`
	ProductID        string         `json:"product_id"`
	RiskRate         string         `json:"risk_rate"`
	CurrencyLTC      MarginCurrency `json:"currency:LTC,omitempty"`
	CurrencyETH      MarginCurrency `json:"currency:ETH,omitempty"`
	CurrencyETC      MarginCurrency `json:"currency:ETC,omitempty"`
	CurrencyBCH      MarginCurrency `json:"currency:BCH,omitempty"`
	CurrencyEOS      MarginCurrency `json:"currency:EOS,omitempty"`
	CurrencyXRP      MarginCurrency `json:"currency:XRP,omitempty"`
	CurrencyTRX      MarginCurrency `json:"currency:TRX,omitempty"`
	CurrencyBSV      MarginCurrency `json:"currency:BSV,omitempty"`
	CurrencyDASH     MarginCurrency `json:"currency:DASH,omitempty"`
	CurrencyNEO      MarginCurrency `json:"currency:NEO,omitempty"`
	CurrencyQTUM     MarginCurrency `json:"currency:QTUM,omitempty"`
	CurrencyIOST     MarginCurrency `json:"currency:IOST,omitempty"`
}

type GetMarginAccountsResult []GetMarginAccountsItem

type MarginNewOrderResult struct {
	ClientOid    string `json:"client_oid"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	OrderID      string `json:"order_id"`
	Result       bool   `json:"result"`
}

type MarginGetOrderResult struct {
	ClientOid      string       `json:"client_oid"`
	CreatedAt      time.Time    `json:"created_at"`
	FilledNotional sjson.Number `json:"filled_notional"`
	FilledSize     sjson.Number `json:"filled_size"`
	Funds          string       `json:"funds"`
	InstrumentID   string       `json:"instrument_id"`
	Notional       string       `json:"notional"`
	OrderID        string       `json:"order_id"`
	OrderType      sjson.Number `json:"order_type"` /*int*/
	Price          sjson.Number `json:"price"`
	PriceAvg       sjson.Number `json:"price_avg"`
	ProductID      string       `json:"product_id"`
	Side           string       `json:"side"`
	Size           sjson.Number `json:"size"`
	Status         string       `json:"status"` // status为state旧版参数，会短期兼容，建议尽早切换state
	State          sjson.Number `json:"state"`  /*int*/ // -2:失败 -1:撤单成功 0:等待成交 1:部分成交 2:完全成交 3:下单中 4:撤单中
	Timestamp      time.Time    `json:"timestamp"`
	Type           string       `json:"type"`
}

type FillItem struct {
	CreatedAt    string  `json:"created_at"`
	ExecType     string  `json:"exec_type"`
	Fee          float64 `json:"fee,string"`
	InstrumentID string  `json:"instrument_id"`
	LedgerID     string  `json:"ledger_id"`
	Liquidity    string  `json:"liquidity"`
	OrderID      string  `json:"order_id"`
	Price        string  `json:"price"`
	ProductID    string  `json:"product_id"`
	Side         string  `json:"side"`
	Size         float64 `json:"size,string"`
	Timestamp    string  `json:"timestamp"`
}

func (r *MarginGetOrderResult) GetState() int64 {
	i, _ := r.State.Int64()
	return i
}
