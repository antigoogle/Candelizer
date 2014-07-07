package main

import (
	"strconv"
)

type CandleChannel struct {
	Exchange   string
	Symbol     string
	Width      int
	Channel    <-chan Candle
	HotChannel <-chan Candle
}

func (cc CandleChannel) String() string {
	return "CandleChannel[Exchange: " + cc.Exchange + ", Symbol: " + cc.Symbol + "]"
}

type Candle struct {
	Date       int
	Open       float64
	Low        float64
	High       float64
	Close      float64
	Volume     float64
	Indicators map[string]float64
}

func (c Candle) String() string {
	return "Candle[Date:" + strconv.Itoa(c.Date) + ", OLHC: " + strconv.FormatFloat(c.Open, 'f', 8, 64) + ", " + strconv.FormatFloat(c.Low, 'f', 8, 64) + ", " + strconv.FormatFloat(c.High, 'f', 8, 64) + ", " + strconv.FormatFloat(c.Close, 'f', 8, 64) + ", Volume: " + strconv.FormatFloat(c.Volume, 'f', 8, 64) + "]"
}
