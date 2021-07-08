package main

import (
	"FakeDataSender/packages/atm"
	"FakeDataSender/packages/ticker"
	"flag"
		"log"
	"time"
)




func init() {

}

func main() {

	var timeout int
	var path string

	flag.IntVar(&timeout, "timeout", 1, "work timeout in minutes")
	flag.StringVar(&path, "path", "./", "output path")

	flag.Parse()


	waitTime :=  time.Minute * time.Duration(timeout)


	// init data for ATM 1
	atm1 := atm.NewATM()
	_ = atm1.SetATMNumber(123456789)
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

	// init data for ATM 2
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

	// init data for ATM 3
	atm3 := atm.NewATM()
	_ = atm3.SetATMNumber(123404321)
	atm3.CashOut[0].DefaultDenominationValue = 50
	atm3.CashOut[1].DefaultDenominationValue = 100
	atm3.CashOut[2].DefaultDenominationValue = 1000
	atm3.CashOut[3].DefaultDenominationValue = 5000

	atm3.CashOut[0].MaximumLoadValue = 500
	atm3.CashOut[1].MaximumLoadValue = 500
	atm3.CashOut[2].MaximumLoadValue = 300
	atm3.CashOut[3].MaximumLoadValue = 300

	for i := range atm3.CashOut {
		atm3.CashOut[i].Load()
	}

// add ticker for all ATMs

	t1 := ticker.NewTicker(atm1)
	t2 := ticker.NewTicker(atm2)
	t3 := ticker.NewTicker(atm3)
	go t1.Run(path)
	go t2.Run(path)
	go t3.Run(path)


	// work time
	time.Sleep(waitTime)

	t1.Ticker.Stop()
	t2.Ticker.Stop()
	t3.Ticker.Stop()
	log.Println("all ATMs work stopped")

}

/*
init ATMs data
set values for denominations and loaded in cassettes
make ticker for values variations

generate data
send data
check response
*/
