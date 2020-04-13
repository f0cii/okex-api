package okex

import (
	"fmt"
	"github.com/MauriceGit/skiplist"
	"strconv"
)

const (
	ActionDepthL2Partial = "partial"
	ActionDepthL2Update  = "update"
)

type Item struct {
	Price  float64
	Amount float64
}

func (e Item) ExtractKey() float64 {
	return e.Price
}

func (e Item) String() string {
	return fmt.Sprintf("%.2f", e.Price)
}

type OrderBook struct {
	InstrumentID string `json:"instrument_id"`
	Asks         []Item `json:"asks"`
	Bids         []Item `json:"bids"`
}

type DepthOrderBook struct {
	instrumentID string // BTC-USD-SWAP
	asks         skiplist.SkipList
	bids         skiplist.SkipList
}

func (d *DepthOrderBook) GetInstrumentID() string {
	return d.instrumentID
}

func (d *DepthOrderBook) Update(action string, data *WSDepthL2Tbt) {
	if action == ActionDepthL2Partial {
		d.asks = skiplist.New()
		d.bids = skiplist.New()
		// 举例: ["411.8", "10", "1", "4"]
		// 411.8为深度价格，10为此价格的合约张数，1为此价格的强平单个数，4为此价格的订单个数。
		for _, ask := range data.Asks {
			price, _ := strconv.ParseFloat(ask[0], 64)
			amount, _ := strconv.ParseFloat(ask[1], 64)
			d.asks.Insert(Item{
				Price:  price,
				Amount: amount,
			})
		}
		for _, bid := range data.Bids {
			price, _ := strconv.ParseFloat(bid[0], 64)
			amount, _ := strconv.ParseFloat(bid[1], 64)
			d.bids.Insert(Item{
				Price:  price,
				Amount: amount,
			})
		}
		return
	}

	if action == ActionDepthL2Update {
		for _, ask := range data.Asks {
			price, _ := strconv.ParseFloat(ask[0], 64)
			amount, _ := strconv.ParseFloat(ask[1], 64)
			if amount == 0 {
				d.asks.Delete(Item{
					Price:  price,
					Amount: amount,
				})
			} else {
				item := Item{
					Price:  price,
					Amount: amount,
				}
				elem, ok := d.asks.Find(item)
				if ok {
					d.asks.ChangeValue(elem, item)
				} else {
					d.asks.Insert(item)
				}
			}
		}
		for _, bid := range data.Bids {
			price, _ := strconv.ParseFloat(bid[0], 64)
			amount, _ := strconv.ParseFloat(bid[1], 64)
			if amount == 0 {
				d.bids.Delete(Item{
					Price:  price,
					Amount: amount,
				})
			} else {
				item := Item{
					Price:  price,
					Amount: amount,
				}
				elem, ok := d.bids.Find(item)
				if ok {
					d.bids.ChangeValue(elem, item)
				} else {
					d.bids.Insert(item)
				}
			}
		}
	}
}

func (d *DepthOrderBook) GetOrderBook(depth int) (result OrderBook) {
	result.InstrumentID = d.instrumentID
	smallest := d.asks.GetSmallestNode()
	if smallest != nil {
		result.Asks = append(result.Asks, smallest.GetValue().(Item))
		count := 1
		node := smallest
		for count < depth {
			node = d.asks.Next(node)
			if node == nil {
				break
			}
			result.Asks = append(result.Asks, node.GetValue().(Item))
			count++
		}
	}

	largest := d.bids.GetLargestNode()
	if largest != nil {
		result.Bids = append(result.Bids, largest.GetValue().(Item))
		count := 1
		node := largest
		for count < depth {
			node = d.bids.Prev(node)
			if node == nil {
				break
			}
			result.Bids = append(result.Bids, node.GetValue().(Item))
			count++
		}
	}
	return
}

func NewDepthOrderBook(instrumentID string) *DepthOrderBook {
	return &DepthOrderBook{
		instrumentID: instrumentID,
		asks:         skiplist.New(),
		bids:         skiplist.New(),
	}
}
