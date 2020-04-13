package okex

import "time"

type WSSwapPositionResult struct {
	Table string               `json:"table"`
	Data  []WSSwapPositionData `json:"data"`
}

type WSSwapPositionHolding struct {
	AvailPosition    string    `json:"avail_position"`
	AvgCost          string    `json:"avg_cost"`
	Last             string    `json:"last"`
	Leverage         string    `json:"leverage"`
	LiquidationPrice string    `json:"liquidation_price"`
	MaintMarginRatio string    `json:"maint_margin_ratio"`
	Margin           string    `json:"margin"`
	Position         string    `json:"position"`
	RealizedPnl      string    `json:"realized_pnl"`
	SettledPnl       string    `json:"settled_pnl"`
	SettlementPrice  string    `json:"settlement_price"`
	Side             string    `json:"side"`
	Timestamp        time.Time `json:"timestamp"`
}

type WSSwapPositionData struct {
	Holding      []WSSwapPositionHolding `json:"holding"`
	InstrumentID string                  `json:"instrument_id"`
	MarginMode   string                  `json:"margin_mode"`
	Timestamp    time.Time               `json:"timestamp"`
}
