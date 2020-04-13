package okex

import "time"

type WSFuturesTicker struct {
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

type WSFuturesTickerResult struct {
	Table string            `json:"table"`
	Data  []WSFuturesTicker `json:"data"`
}

type WSFuturesTrade struct {
	Side         string    `json:"side"`
	TradeID      string    `json:"trade_id"`
	Price        string    `json:"price"`
	Qty          string    `json:"qty"`
	InstrumentID string    `json:"instrument_id"`
	Timestamp    time.Time `json:"timestamp"`
}

type WSFuturesTradeResult struct {
	Table string           `json:"table"`
	Data  []WSFuturesTrade `json:"data"`
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

type WSFuturesAccountResult struct {
	Table string                 `json:"table"`
	Data  []WSFuturesAccountData `json:"data"`
}

type WSFuturesAccount struct {
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

type WSFuturesAccountData struct {
	BTC *WSFuturesAccount `json:"BTC"`
	LTC *WSFuturesAccount `json:"LTC"`
	ETH *WSFuturesAccount `json:"ETH"`
	ETC *WSFuturesAccount `json:"ETC"`
	XRP *WSFuturesAccount `json:"XRP"`
	EOS *WSFuturesAccount `json:"EOS"`
	BCH *WSFuturesAccount `json:"BCH"`
	BSV *WSFuturesAccount `json:"BSV"`
	TRX *WSFuturesAccount `json:"TRX"`
}

type WSFuturesPositionResult struct {
	Table string              `json:"table"`
	Data  []WSFuturesPosition `json:"data"`
}

type WSFuturesPosition struct {
	LongQty               string    `json:"long_qty"`
	LongAvailQty          string    `json:"long_avail_qty"`
	LongMargin            string    `json:"long_margin"`
	LongLiquiPrice        string    `json:"long_liqui_price"`
	LongPnlRatio          string    `json:"long_pnl_ratio"`
	LongAvgCost           string    `json:"long_avg_cost"`
	LongSettlementPrice   string    `json:"long_settlement_price"`
	RealisedPnl           string    `json:"realised_pnl"`
	ShortQty              string    `json:"short_qty"`
	ShortAvailQty         string    `json:"short_avail_qty"`
	ShortMargin           string    `json:"short_margin"`
	ShortLiquiPrice       string    `json:"short_liqui_price"`
	ShortPnlRatio         string    `json:"short_pnl_ratio"`
	ShortAvgCost          string    `json:"short_avg_cost"`
	ShortSettlementPrice  string    `json:"short_settlement_price"`
	InstrumentID          string    `json:"instrument_id"`
	LongLeverage          string    `json:"long_leverage"`
	ShortLeverage         string    `json:"short_leverage"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	Timestamp             time.Time `json:"timestamp"`
	MarginMode            string    `json:"margin_mode"`
	ShortMarginRatio      string    `json:"short_margin_ratio"`
	ShortMaintMarginRatio string    `json:"short_maint_margin_ratio"`
	ShortPnl              string    `json:"short_pnl"`
	ShortUnrealisedPnl    string    `json:"short_unrealised_pnl"`
	LongMarginRatio       string    `json:"long_margin_ratio"`
	LongMaintMarginRatio  string    `json:"long_maint_margin_ratio"`
	LongPnl               string    `json:"long_pnl"`
	LongUnrealisedPnl     string    `json:"long_unrealised_pnl"`
	LongOpenOutstanding   string    `json:"long_open_outstanding"`
	ShortOpenOutstanding  string    `json:"short_open_outstanding"`
	LongSettledPnl        string    `json:"long_settled_pnl"`
	ShortSettledPnl       string    `json:"short_settled_pnl"`
	Last                  string    `json:"last"`
}

type WSFuturesOrder struct {
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

type WSFuturesOrderResult struct {
	Table string           `json:"table"`
	Data  []WSFuturesOrder `json:"data"`
}
