package main

import (
	"strconv"
)

type TradeChannel struct {
	Exchange string
	Symbol   string
	Channel  <-chan Trade
}

func (tc TradeChannel) String() string {
	return "TradeChannel[Exchange: " + tc.Exchange + ", Symbol: " + tc.Symbol + "]"
}

type Trade struct {
	Date   int
	Price  float64
	Amount float64
}

func (t Trade) String() string {
	return "Trade[Date: " + strconv.Itoa(t.Date) + ", Price: " + strconv.FormatFloat(t.Price, 'f', 8, 64) + ", Amount: " + strconv.FormatFloat(t.Amount, 'f', 8, 64) + "]"
}
