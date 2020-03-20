package okex

import "testing"

func TestGetIndexConstituents(t *testing.T) {
	c := NewTestClient()
	result, err := c.GetIndexConstituents("BTC-USDT")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", result)
}
