package okex

import "testing"

func TestGetFinancialRates(t *testing.T) {
	c := NewTestClient()
	result, err := c.GetFinancialRates()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", result)
}
