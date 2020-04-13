package okex

import "time"

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
