package ticker

import (
	"FakeDataSender/packages/atm"
	"FakeDataSender/packages/tojson"
	"fmt"
	"time"
)

const tickTimeout = 100

type Ticker struct {
	tickTimeout int
	Ticker      *time.Ticker
	atm         *atm.ATM
}

func NewTicker(a *atm.ATM) *Ticker {
	t := new(Ticker)
	t.atm = a
	t.Ticker = time.NewTicker(time.Millisecond * tickTimeout)
	return t
}

func (t *Ticker) Run() {
	go func() {
		for ti := range t.Ticker.C {
			fmt.Println("Tick at", ti)

			// ATM operations
			t.atm.Dice()
			fmt.Println(tojson.PackToJSON(t.atm))
		}
	}()
}
