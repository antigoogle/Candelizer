package types

import (
	"sync"
)

type CS struct {
	Data   []Candle
	Update <-chan Candle
	sync.Mutex
}
