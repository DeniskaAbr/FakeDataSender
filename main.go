package main

import (
	"FakeDataSender/packages/atm"
	"FakeDataSender/packages/ticker"
	"FakeDataSender/packages/tojson"
	"fmt"
	"time"
)

const tickTask = time.Hour * 1

func init() {

}

func main() {

	atm1 := atm.NewATM()
	err := atm1.SetATMNumber(123456789)
	atm1.CashOut[0].DefaultDenominationValue = 50
	atm1.CashOut[1].DefaultDenominationValue = 100
	atm1.CashOut[2].DefaultDenominationValue = 1000
	atm1.CashOut[3].DefaultDenominationValue = 5000

	atm1.CashOut[0].MaximumLoadValue = 500
	atm1.CashOut[1].MaximumLoadValue = 500
	atm1.CashOut[2].MaximumLoadValue = 300
	atm1.CashOut[3].MaximumLoadValue = 300

	for i := range atm1.CashOut {
		atm1.CashOut[i].Load()
	}

	fmt.Println(atm1)

	err = atm1.Dispense(150)
	fmt.Println(err)
	err = atm1.Dispense(150)
	fmt.Println(err)

	fmt.Println(atm1)

	str := tojson.PackToJSON(atm1)
	if len(str) != 0 {
		fmt.Println(str)
	}

	atm2 := atm.NewATM()
	_ = atm2.SetATMNumber(987654321)
	atm2.CashOut[0].DefaultDenominationValue = 50
	atm2.CashOut[1].DefaultDenominationValue = 100
	atm2.CashOut[2].DefaultDenominationValue = 1000
	atm2.CashOut[3].DefaultDenominationValue = 5000

	atm2.CashOut[0].MaximumLoadValue = 500
	atm2.CashOut[1].MaximumLoadValue = 500
	atm2.CashOut[2].MaximumLoadValue = 300
	atm2.CashOut[3].MaximumLoadValue = 300

	for i := range atm2.CashOut {
		atm2.CashOut[i].Load()
	}
	atm2.CashOut[0].Loaded = 3

	t1 := ticker.NewTicker(atm1)
	t2 := ticker.NewTicker(atm2)
	go t2.Run()
	go t1.Run()

	time.Sleep(tickTask)
	fmt.Println("Просим тикер остановиться")
	t1.Ticker.Stop()
	t1.Ticker.Stop()
	fmt.Println("Ticker stopped")

	fmt.Println("вышли в основной код")
	fmt.Println(atm1.CashOut[0].Loaded)
	fmt.Println(atm2.CashOut[0].Loaded)
}

/*
init ATMs data
set values for denominations and loaded in cassettes
make ticker for values variations

generate data
send data
check response
*/
