package indicators

import (
	"fmt"
	"github.com/flower1024/Candelizer/types"
)

type Indicators struct {
	cs     types.CS
	Update <-chan types.Candle
}

func (i Indicators) Run(indicator map[string]interface{}, hotCandle types.Candle) {
	i.cs.Lock()
	defer i.cs.Unlock()
	switch indicator["Indicator"] {
	case "ema":
		EMA(indicator["ema"].(float64), indicator["field"].(string), i.cs.Data, hotCandle)
	default:
		fmt.Printf("Unkown Indicator: %s", indicator["Indicator"])
	}
}
