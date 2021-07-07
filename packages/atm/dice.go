package atm

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

func (atm *ATM) Dice() {
	var info string
	for i, cassette := range atm.CashOut {
		if cassette.MaximumLoadValue > 0 && cassette.Status == "OK" {
			loadPercent := (float64(cassette.Loaded) / float64(cassette.MaximumLoadValue)) * 100.0
			if loadPercent < 25.0 {
				atm.CashOut[i].StatusSwitch()
				info = "cassette error"
				s := fmt.Sprintf(`%v cassette %v load low %g percent`, atm.AtmNumber, i+1, loadPercent)
				log.Println(s)
			}
		}
		if cassette.Status == "ERROR" && atm.SafetyDevice.SafeDoor == "closed" && len(info) == 0 {

			atm.SafetyDevice.SafeDoor = "open"
			info = "safe open"
			s := fmt.Sprintf("%v safe open ", atm.AtmNumber)
			log.Println(s)
		}
		if cassette.Status == "ERROR" && atm.SafetyDevice.SafeDoor == "open" && len(info) == 0 {
			for i := range atm.CashOut {
				atm.CashOut[i].Load()
			}

			info = "atm loaded"
			s := fmt.Sprintf("%v all cassettes reloaded", atm.AtmNumber)
			log.Println(s)

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
		atm.ResetAllErrors()
		s := fmt.Sprintf("%v safe closed", atm.AtmNumber)
		log.Println(s)

	default:
		if !atm.Error() {
			randomCash := rand.Intn(500)
			atm.Dispense(randomCash * cassettesInfo[len(cassettesInfo)-1].Denomination)
			summa := randomCash * cassettesInfo[len(cassettesInfo)-1].Denomination
			s := fmt.Sprintf("%v dispence %v", atm.AtmNumber, summa)
			log.Println(s)
		}
	}
}
