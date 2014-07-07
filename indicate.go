package main

import (
	"fmt"
)

type Indicators struct {
	cs     CS
	Update <-chan Candle
}

func (i Indicators) Request() []Candle {
	return i.cs.Request()
}

func Indicate(cs CS) Indicators {
	fmt.Printf("INIT Indicators\n")

	update := make(chan Candle)
	i := Indicators{Update: update, cs: cs}

	go func() {
		for {
			hotCandle := <-i.cs.Update
			update <- hotCandle
		}
	}()

	return i
}
