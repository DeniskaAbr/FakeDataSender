package atm

import (
	"fmt"
	"math/rand"
	"sort"
)

func (atm *ATM) Dice() {
	var info string
	for i, cassette := range atm.CashOut {
		if cassette.MaximumLoadValue > 0 && cassette.Status == "OK" {
			if (float64(cassette.Loaded)/float64(cassette.MaximumLoadValue))*100.0 < 25.0 {
				fmt.Println(">>>>>>>>")
				fmt.Println("cassette low")
				fmt.Println("<<<<<<<<")
				atm.CashOut[i].StatusSwitch()
				info = "cassette error"
			}
		}
		if cassette.Status == "ERROR" && atm.SafetyDevice.SafeDoor == "closed" && len(info) == 0 {
			fmt.Println(">>>>>>>>")
			fmt.Println("open safe")
			fmt.Println("<<<<<<<<")
			atm.SafetyDevice.SafeDoor = "open"
			info = "safe open"
		}
		if cassette.Status == "ERROR" && atm.SafetyDevice.SafeDoor == "open" && len(info) == 0 {
			for i := range atm.CashOut {
				atm.CashOut[i].Load()
			}
			fmt.Println(">>>>>>>>")
			fmt.Println("load all cassettes")
			fmt.Println("<<<<<<<<")
			atm.ResetAllErrors()
			info = "atm loaded"
		}
	}

	var cassettesInfo []Cassette
	cassettesInfo = append(cassettesInfo, atm.CashOut...)
	for i := range atm.CashOut {
		atm.CashOut[i].CassetteIndex = i
	}
	// sort by denomination values
	sort.Slice(cassettesInfo, func(i, j int) bool {
		return cassettesInfo[i].Denomination > cassettesInfo[j].Denomination
	})

	switch info {
	case "cassette error":
	case "safe open":
	case "atm loaded":
		atm.SafetyDevice.SafeDoor = "closed"
		fmt.Println(">>>>>>>>")
		fmt.Println("close safe")
		fmt.Println("<<<<<<<<")

	default:
		if !atm.Error() {
			randomCash := rand.Intn(500)
			atm.Dispense(randomCash * cassettesInfo[len(cassettesInfo)-1].Denomination)
			fmt.Println(">>>>>>>>")
			fmt.Println("Dispence")
			fmt.Println(randomCash * cassettesInfo[len(cassettesInfo)-1].Denomination)
			fmt.Println("<<<<<<<<")
		}
	}
}
