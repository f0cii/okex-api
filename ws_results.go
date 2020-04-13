package okex

import "time"

type WSTicker struct {
	Last           string    `json:"last"`
	Open24H        string    `json:"open_24h"`
	BestBid        string    `json:"best_bid"`
	High24H        string    `json:"high_24h"`
	Low24H         string    `json:"low_24h"`
	Volume24H      string    `json:"volume_24h"`
	VolumeToken24H string    `json:"volume_token_24h"`
	BestAsk        string    `json:"best_ask"`
	OpenInterest   string    `json:"open_interest"`
	InstrumentID   string    `json:"instrument_id"`
	Timestamp      time.Time `json:"timestamp"`
	BestBidSize    string    `json:"best_bid_size"`
	BestAskSize    string    `json:"best_ask_size"`
	LastQty        string    `json:"last_qty"`
}

type WSTickerResult struct {
	Table string     `json:"table"`
	Data  []WSTicker `json:"data"`
}

type WSTrade struct {
	Side         string    `json:"side"`
	TradeID      string    `json:"trade_id"`
	Price        string    `json:"price"`
	Qty          string    `json:"qty"`
	InstrumentID string    `json:"instrument_id"`
	Timestamp    time.Time `json:"timestamp"`
}

type WSTradeResult struct {
	Table string    `json:"table"`
	Data  []WSTrade `json:"data"`
}

type WSDepthL2Tbt struct {
	InstrumentID string     `json:"instrument_id"`
	Asks         [][]string `json:"asks"`
	Bids         [][]string `json:"bids"`
	Timestamp    time.Time  `json:"timestamp"`
	Checksum     int        `json:"checksum"`
}

type WSDepthL2TbtResult struct {
	Table  string         `json:"table"`
	Action string         `json:"action"`
	Data   []WSDepthL2Tbt `json:"data"`
}

type WSAccountResult struct {
	Table string          `json:"table"`
	Data  []WSAccountData `json:"data"`
}

type WSAccount struct {
	Available         string    `json:"available"`
	CanWithdraw       string    `json:"can_withdraw"`
	Currency          string    `json:"currency"`
	Equity            string    `json:"equity"`
	LiquiMode         string    `json:"liqui_mode"`
	MaintMarginRatio  string    `json:"maint_margin_ratio"`
	Margin            string    `json:"margin"`
	MarginForUnfilled string    `json:"margin_for_unfilled"`
	MarginFrozen      string    `json:"margin_frozen"`
	MarginMode        string    `json:"margin_mode"`
	MarginRatio       string    `json:"margin_ratio"`
	OpenMax           string    `json:"open_max"`
	RealizedPnl       string    `json:"realized_pnl"`
	Timestamp         time.Time `json:"timestamp"`
	TotalAvailBalance string    `json:"total_avail_balance"`
	Underlying        string    `json:"underlying"`
	UnrealizedPnl     string    `json:"unrealized_pnl"`
}

type WSAccountData struct {
	BTC *WSAccount `json:"BTC"`
	LTC *WSAccount `json:"LTC"`
	ETH *WSAccount `json:"ETH"`
	ETC *WSAccount `json:"ETC"`
	XRP *WSAccount `json:"XRP"`
	EOS *WSAccount `json:"EOS"`
	BCH *WSAccount `json:"BCH"`
	BSV *WSAccount `json:"BSV"`
	TRX *WSAccount `json:"TRX"`
}

type WSOrder struct {
	Leverage     string    `json:"leverage"`
	LastFillTime time.Time `json:"last_fill_time"`
	FilledQty    string    `json:"filled_qty"`
	Fee          string    `json:"fee"`
	PriceAvg     string    `json:"price_avg"`
	Type         string    `json:"type"`
	ClientOid    string    `json:"client_oid"`
	LastFillQty  string    `json:"last_fill_qty"`
	InstrumentID string    `json:"instrument_id"`
	LastFillPx   string    `json:"last_fill_px"`
	Pnl          string    `json:"pnl"`
	Size         string    `json:"size"`
	Price        string    `json:"price"`
	LastFillID   string    `json:"last_fill_id"`
	ErrorCode    string    `json:"error_code"`
	State        string    `json:"state"`
	ContractVal  string    `json:"contract_val"`
	OrderID      string    `json:"order_id"`
	OrderType    string    `json:"order_type"`
	Timestamp    time.Time `json:"timestamp"`
	Status       string    `json:"status"`
}

type WSOrderResult struct {
	Table string    `json:"table"`
	Data  []WSOrder `json:"data"`
}
