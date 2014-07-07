package candles

import (
	"fmt"
	"github.com/flower1024/Candelizer/types"
)

type Indicators struct {
	cs     CS
	Update <-chan types.Candle
}

func (i Indicators) Request() []types.Candle {
	return i.cs.Request()
}

func Indicate(cs CS) Indicators {
	fmt.Printf("INIT Indicators\n")

	update := make(chan types.Candle)
	i := Indicators{Update: update, cs: cs}

	go func() {
		for {
			hotCandle := <-i.cs.Update
			update <- hotCandle
		}
	}()

	return i
}
