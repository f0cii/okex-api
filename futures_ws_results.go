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
