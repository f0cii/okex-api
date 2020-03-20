package okex

import "time"

type GetMarginAccountsByInstrumentItem struct {
	Available   float64 `json:"available,string"`    // 可用于交易的数量
	Balance     float64 `json:"balance,string"`      // 余额
	Borrowed    float64 `json:"borrowed,string"`     // 已借币（已借未还的部分）
	CanWithdraw float64 `json:"can_withdraw,string"` // 可划转数量
	Frozen      float64 `json:"frozen,string"`
	Hold        float64 `json:"hold,string"` // 冻结（不可用）
	Holds       float64 `json:"holds,string"`
	LendingFee  float64 `json:"lending_fee,string"` // 利息（未还的利息）
}

type GetMarginAccountsByInstrumentResult struct {
	CurrencyBTC      GetMarginAccountsByInstrumentItem `json:"currency:BTC"`
	CurrencyLTC      GetMarginAccountsByInstrumentItem `json:"currency:LTC"`
	CurrencyETH      GetMarginAccountsByInstrumentItem `json:"currency:ETH"`
	CurrencyETC      GetMarginAccountsByInstrumentItem `json:"currency:ETC"`
	CurrencyBCH      GetMarginAccountsByInstrumentItem `json:"currency:BCH"`
	CurrencyEOS      GetMarginAccountsByInstrumentItem `json:"currency:EOS"`
	CurrencyXRP      GetMarginAccountsByInstrumentItem `json:"currency:XRP"`
	CurrencyUSDT     GetMarginAccountsByInstrumentItem `json:"currency:USDT"`
	LiquidationPrice float64                           `json:"liquidation_price,string"`
	MarginRatio      string                            `json:"margin_ratio"`
	RiskRate         string                            `json:"risk_rate"`
}

type PostMarginAccountsBorrowResult struct {
	BorrowID  string `json:"borrow_id"`
	ClientOid string `json:"client_oid"`
	Result    bool   `json:"result"`
}

type GetMarginAccountsBorrowedByInstrumentIdItem struct {
	Amount           float64   `json:"amount,string"`
	BorrowID         string    `json:"borrow_id"`
	CreatedAt        time.Time `json:"created_at"`
	Currency         string    `json:"currency"`
	ForceRepayTime   time.Time `json:"force_repay_time"`
	InstrumentID     string    `json:"instrument_id"`
	Interest         float64   `json:"interest,string"`
	LastInterestTime time.Time `json:"last_interest_time"`
	PaidInterest     float64   `json:"paid_interest,string"`
	ProductID        string    `json:"product_id"`
	Rate             float64   `json:"rate,string"`
	RepayAmount      float64   `json:"repay_amount,string"`
	RepayInterest    float64   `json:"repay_interest,string"`
	ReturnedAmount   float64   `json:"returned_amount,string"`
	Timestamp        time.Time `json:"timestamp"`
}

type GetMarginAccountsBorrowedByInstrumentIdResult []GetMarginAccountsBorrowedByInstrumentIdItem

type GetMarginAccountsAvailabilityByInstrumentIdItem struct {
	Available     float64 `json:"available,string"` // 当前最大可借
	Leverage      float64 `json:"leverage,string"`  // 最大杠杆倍数
	LeverageRatio float64 `json:"leverage_ratio,string"`
	Rate          float64 `json:"rate,string"` // 借币利率
}

type GetMarginAccountsAvailabilityByInstrumentIdResult []struct {
	CurrencyBTC GetMarginAccountsAvailabilityByInstrumentIdItem `json:"currency:BTC"`
	CurrencyLTC GetMarginAccountsAvailabilityByInstrumentIdItem `json:"currency:LTC"`
	CurrencyETH GetMarginAccountsAvailabilityByInstrumentIdItem `json:"currency:ETH"`
	CurrencyETC GetMarginAccountsAvailabilityByInstrumentIdItem `json:"currency:ETC"`
	CurrencyBCH GetMarginAccountsAvailabilityByInstrumentIdItem `json:"currency:BCH"`
	CurrencyEOS GetMarginAccountsAvailabilityByInstrumentIdItem `json:"currency:EOS"`
	//CurrencyXRP GetMarginAccountsAvailabilityByInstrumentIdItem `json:"currency:XRP"`
	CurrencyUSDT GetMarginAccountsAvailabilityByInstrumentIdItem `json:"currency:USDT"`
	InstrumentID string                                          `json:"instrument_id"`
	ProductID    string                                          `json:"product_id"`
}

type PostMarginAccountsRepaymentResult struct {
	ClientOid   string `json:"client_oid"`
	RepaymentID string `json:"repayment_id"`
	Result      bool   `json:"result"`
}
