package okex

/*
获取平台所有币种列表。并非所有币种都可被用于交易。在ISO 4217标准中未被定义的币种代码可能使用的是自定义代码。

HTTP请求
GET /api/account/v3/currencies

*/
func (client *Client) GetAccountCurrencies() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, _, err := client.Request(GET, ACCOUNT_CURRENCIES, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
钱包账户信息
获取钱包账户所有资产列表，查询各币种的余额、冻结和可用等信息。

HTTP请求
GET /api/account/v3/wallet
*/
func (client *Client) GetAccountWallet() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, _, err := client.Request(GET, ACCOUNT_WALLET, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
单一币种账户信息
获取钱包账户单个币种的余额、冻结和可用等信息。

HTTP请求
GET /api/account/v3/wallet/<currency>

请求示例
GET /api/account/v3/wallet/btc
*/
func (client *Client) GetAccountWalletByCurrency(currency string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := GetCurrencyUri(ACCOUNT_WALLET_CURRENCY, currency)

	if _, _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
单一币种账户信息
获取钱包账户单个币种的余额、冻结和可用等信息。

HTTP请求
GET /api/account/v3/wallet/<currency>
*/
func (client *Client) GetAccountWithdrawalFeeByCurrency(currency *string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := ACCOUNT_WITHRAWAL_FEE
	if currency != nil && len(*currency) > 0 {
		params := NewParams()
		params["currency"] = *currency
		uri = BuildParams(uri, params)
	}

	if _, _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
查询最近所有币种的提币记录

HTTP请求
GET /api/account/v3/withdrawal/history
*/
func (client *Client) GetAccountWithdrawalHistory() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, _, err := client.Request(GET, ACCOUNT_WITHRAWAL_HISTORY, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
查询单个币种的提币记录。

HTTP请求
GET /api/account/v3/withdrawal/history/<currency>
*/
func (client *Client) GetAccountWithdrawalHistoryByCurrency(currency string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := GetCurrencyUri(ACCOUNT_WITHRAWAL_HISTORY_CURRENCY, currency)

	if _, _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取充值地址
获取各个币种的充值地址，包括曾使用过的老地址。

HTTP请求
GET /api/account/v3/deposit/address

请求示例
GET /api/account/v3/deposit/address?currency=btc
*/
func (client *Client) GetAccountDepositAddress(currency string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}
	params := NewParams()
	params["currency"] = currency

	uri := BuildParams(ACCOUNT_DEPOSIT_ADDRESS, params)

	if _, _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取所有币种充值记录
获取所有币种的充值记录。为最近一百条数据

HTTP请求
GET /api/account/v3/deposit/history
*/
func (client *Client) GetAccountDepositHistory() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, _, err := client.Request(GET, ACCOUNT_DEPOSIT_HISTORY, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取单个币种充值记录
获取单个币种的充值记录，为最近一百条数据

HTTP
GET /api/account/v3/deposit/history/<currency>
*/
func (client *Client) GetAccountDepositHistoryByCurrency(currency string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := GetCurrencyUri(ACCOUNT_DEPOSIT_HISTORY_CURRENCY, currency)

	if _, _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
账单流水查询
查询钱包账户账单流水。流水会分页，并且按时间倒序排序和储存，最新的排在最前面。请参阅分页部分以获取第一页之后的其他记录。为最近三个月的数据

HTTP请求
GET /api/account/v3/ledger

请求示例
GET /api/account/v3/ledger?type=2&currency=btc&from=4&limit=10
*/
func (client *Client) GetAccountLeger(optionalParams *map[string]string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}
	uri := ACCOUNT_LEDGER
	if optionalParams != nil && len(*optionalParams) > 0 {
		uri = BuildParams(uri, *optionalParams)
	}

	if _, _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
提币
提币到OKCoin国际站账户，OKEx账户或数字货币地址。

HTTP请求
POST /api/account/v3/withdrawal
*/
func (client *Client) PostAccountWithdrawal(
	currency, to_address, trade_pwd string, destination int32, amount, fee float32) ([]byte, *map[string]interface{}, error) {

	r := map[string]interface{}{}

	withdrawlInfo := map[string]interface{}{}
	withdrawlInfo["amount"] = amount
	withdrawlInfo["destination"] = destination
	withdrawlInfo["fee"] = fee
	withdrawlInfo["currency"] = currency
	withdrawlInfo["to_address"] = to_address
	withdrawlInfo["trade_pwd"] = trade_pwd

	var respBody []byte
	var err error
	if respBody, _, err = client.Request(GET, ACCOUNT_WITHRAWAL, withdrawlInfo, &r); err != nil {
		return respBody, nil, err
	}

	return respBody, &r, nil
}

/*
资金划转
OKEx站内在钱包账户、交易账户和子账户之间进行资金划转。

限速规则：2次/s
HTTP请求
POST /api/account/v3/transfer
*/
func (client *Client) PostAccountTransfer(
	currency string, from, to int32, amount float32, optionalParams map[string]string) ([]byte, *map[string]interface{}, error) {

	r := map[string]interface{}{}

	transferInfo := map[string]interface{}{}
	transferInfo["currency"] = currency
	transferInfo["from"] = from
	transferInfo["to"] = to
	transferInfo["amount"] = amount

	if optionalParams != nil && len(optionalParams) > 0 {
		if from == 0 || to == 0 {
			transferInfo["sub_account"] = optionalParams["sub_account"]
		}
		transferInfo["instrument_id"] = optionalParams["instrument_id"]
		if to == 5 || to == 3 { // 杠杆转入币对 / usdt保证金合约underlying
			transferInfo["to_instrument_id"] = optionalParams["to_instrument_id"]
		}
	}

	var respBody []byte
	var err error
	if respBody, _, err = client.Request(POST, ACCOUNT_TRANSFER, transferInfo, &r); err != nil {
		return respBody, nil, err
	}

	return respBody, &r, nil
}
