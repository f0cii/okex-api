package okex

import "time"

type GetSpotAccountsResult []GetSpotAccountsResultItem

type GetSpotAccountsResultItem struct {
	Frozen    string  `json:"frozen"`
	Hold      string  `json:"hold"`
	ID        string  `json:"id"`
	Currency  string  `json:"currency"`
	Balance   float64 `json:"balance,string"`
	Available float64 `json:"available,string"`
	Holds     string  `json:"holds"`
}

// price	String	价格
// size	String	数量
// num_orders	String	组成此条深度的订单数量
type SpotInstrumentBookResult struct {
	Asks      [][]string `json:"asks,string"`
	Bids      [][]string `json:"bids,string"`
	Timestamp string     `json:"timestamp"`
}

type SpotNewOrderResult struct {
	ClientOid    string `json:"client_oid"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	OrderID      string `json:"order_id"`
	Result       bool   `json:"result"`
}

type SpotGetOrderResult struct {
	ClientOid      string    `json:"client_oid"`
	CreatedAt      time.Time `json:"created_at"`
	FilledNotional float64   `json:"filled_notional,string"`
	FilledSize     float64   `json:"filled_size,string"`
	Funds          string    `json:"funds"`
	InstrumentID   string    `json:"instrument_id"`
	Notional       string    `json:"notional"`
	OrderID        string    `json:"order_id"`
	OrderType      string    `json:"order_type"`
	Price          string    `json:"price"`
	ProductID      string    `json:"product_id"`
	Side           string    `json:"side"`
	Size           string    `json:"size"`
	Status         string    `json:"status"`
	State          int       `json:"state,string"`
	Timestamp      time.Time `json:"timestamp"`
	Type           string    `json:"type"`
}

type GetSpotAccountsCurrencyResult struct {
	Frozen    string `json:"frozen"`
	Hold      string `json:"hold"`
	ID        string `json:"id"`
	Currency  string `json:"currency"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Holds     string `json:"holds"`
}
