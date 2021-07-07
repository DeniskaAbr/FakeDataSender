package atm

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

const errorStatus = "ERROR"
const okStatus = "OK"

type ATM struct {
	AtmNumber int
	CashOut   []Cassette
	CashIn
	MotorCardReader
	ContactlessCardReader
	EppPinPad
	SafetyDevice
}

type Cassette struct {
	CassetteIndex            int
	MaximumLoadValue         int
	DefaultDenominationValue int
	DefaultCurrencyCode      CurrencyCode

	Loaded       int
	Rejected     int
	Dispensed    int
	Currency     CurrencyCode
	Denomination int
	Status       string
}

type CashIn struct {
	ValidatorFirmwareTemplate string
	Status                    string
	Accepted                  int
}

type MotorCardReader struct {
	Firmware string
	Status   string
}

type ContactlessCardReader struct {
	Firmware string
	Status   string
}

type EppPinPad struct {
	SerialKey string
	Firmware  string
	Status    string
}

type SafetyDevice struct {
	Status   string
	SafeDoor string
}

func (cst *Cassette) Load() {
	// loaded reset loaded to maximum value

	if cst.DefaultCurrencyCode == 0 {
		cst.DefaultCurrencyCode = CurrencyValue.Rouble
	}
	if cst.DefaultDenominationValue == 0 {
		cst.DefaultDenominationValue = 100
	}
	if cst.MaximumLoadValue == 0 {
		cst.MaximumLoadValue = 300

	}

	cst.Loaded = cst.MaximumLoadValue
	cst.Denomination = cst.DefaultDenominationValue
	cst.Currency = cst.DefaultCurrencyCode

	cst.Rejected = 0
	cst.Dispensed = 0
	cst.Status = okStatus
}

func (cst *Cassette) Reject(v int) error {
	if v > cst.Loaded {
		return errors.New("")
	}
	cst.Rejected += v
	cst.Loaded -= v
	return nil
}

func (cst *Cassette) Dispense(v int) error {
	if v > cst.Loaded {
		return errors.New("loaded amount is low")
	}
	cst.Dispensed += v
	cst.Loaded -= v
	return nil
}

func (cst *Cassette) StatusSwitch() {
	if cst.Status == errorStatus {
		cst.Status = okStatus
	} else {
		cst.Status = errorStatus
	}
}

// SetATMNumber Set ATM number by value
func (atm *ATM) SetATMNumber(n int) error {
	if n > 0 && len(strconv.Itoa(n)) < 10 {
		atm.AtmNumber = n
	} else {
		return errors.New("number not set")
	}
	return nil
}

// ResetAllErrors Reset all errors for ATM Devices
func (atm *ATM) ResetAllErrors() {
	for i := range atm.CashOut {
		atm.CashOut[i].Status = okStatus
	}
	atm.CashIn.Status = okStatus
	atm.MotorCardReader.Status = okStatus
	atm.ContactlessCardReader.Status = okStatus
	atm.EppPinPad.Status = okStatus
}

func (atm *ATM) CashInStatusSwitch() {
	if atm.CashIn.Status == errorStatus {
		atm.CashIn.Status = okStatus
	} else {
		atm.CashIn.Status = errorStatus
	}
}

func (atm *ATM) MotorCardReaderStatusSwitch() {
	if atm.MotorCardReader.Status == errorStatus {
		atm.MotorCardReader.Status = okStatus
	} else {
		atm.MotorCardReader.Status = errorStatus
	}
}

func (atm *ATM) AvailableCash() int {
	var cash int
	for _, cassette := range atm.CashOut {
		cash = cassette.Loaded * cassette.Denomination
	}
	return cash
}

func (atm *ATM) Error() bool {
	var b bool

	if atm.CashIn.Status == errorStatus ||
		atm.MotorCardReader.Status == errorStatus ||
		atm.ContactlessCardReader.Status == errorStatus ||
		atm.EppPinPad.Status == errorStatus {
		b = false
	}
	for i := range atm.CashOut {
		if atm.CashOut[i].Status == errorStatus {
			b = false
		}
	}

	return b
}

func (atm *ATM) ContactlessCardReaderStatusSwitch() {
	if atm.ContactlessCardReader.Status == errorStatus {
		atm.ContactlessCardReader.Status = okStatus
	} else {
		atm.ContactlessCardReader.Status = errorStatus
	}
}

// NewATM Create new ATM Instance
func NewATM() *ATM {
	atm := ATM{
		AtmNumber:             0,
		CashOut:               nil,
		CashIn:                CashIn{},
		MotorCardReader:       MotorCardReader{},
		ContactlessCardReader: ContactlessCardReader{},
		EppPinPad:             EppPinPad{},
		SafetyDevice:          SafetyDevice{},
	}

	// load all ATM cassette by default values (300 100 810 0 0 OK)
	// by default ATM have 4 cassette
	atm.CashOut = append(
		atm.CashOut,
		Cassette{},
		Cassette{},
		Cassette{},
		Cassette{})
	for i := range atm.CashOut {
		atm.CashOut[i].Load()
	}

	// Set CashIn by default not null values
	atm.CashIn = CashIn{
		ValidatorFirmwareTemplate: "04.07.123",
		Status:                    "OK",
		Accepted:                  0,
	}

	atm.MotorCardReader = MotorCardReader{
		Firmware: "04.07.123",
		Status:   "OK",
	}
	atm.ContactlessCardReader = ContactlessCardReader{
		Firmware: "04.07.123",
		Status:   "OK",
	}

	atm.EppPinPad = EppPinPad{
		SerialKey: "RE86329TS-skj68273",
		Firmware:  "K98692839",
		Status:    "OK",
	}

	atm.SafetyDevice = SafetyDevice{
		Status:   "блокировка отключена",
		SafeDoor: "closed",
	}

	return &atm
}

func (atm *ATM) Dispense(v int) error {
	// get currencies and available banknotes in cassettes
	var transaction []DispenseTransaction
	for i := range atm.CashOut {
		atm.CashOut[i].CassetteIndex = i
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
	// fmt.Println(cassettesInfo)
	// fmt.Println(atm.CashOut)
	if v%cassettesInfo[len(cassettesInfo)-1].Denomination != 0 {
		// fmt.Println("Not divide to small banknote value")
		return errors.New("cash not divide to denomination value")
	}
	// fmt.Println("Divide to small banknote value")
	var vd = v
	var module int
	for _, cassette := range cassettesInfo {
		module = (vd - (vd % cassette.Denomination)) / cassette.Denomination
		if vd > 0 {
			if cassette.Loaded > 0 {
				if cassette.Loaded < module {
					transaction = append(transaction, DispenseTransaction{
						CassetteIndex:  cassette.CassetteIndex,
						BanknotesCount: cassette.Loaded,
					})
					vd -= cassette.Denomination * cassette.Loaded
				} else {
					transaction = append(transaction, DispenseTransaction{
						CassetteIndex:  cassette.CassetteIndex,
						BanknotesCount: module,
					})
					vd %= cassette.Denomination
				}
			}
		}
	}
	err := makeTransact(vd, v, transaction, atm)
	if err != nil {
		return err
	}
	return nil
}

func makeTransact(vd, v int, transaction []DispenseTransaction, atm *ATM) error {
	switch {
	case vd > 0 && len(transaction) > 0:
		str := fmt.Sprintf("select other exchange value equal %d", v-vd)
		return errors.New(str)
	case len(transaction) == 0:
		str := "no money to dispense"
		return errors.New(str)
	case vd == 0:
		for _, transact := range transaction {
			_ = atm.CashOut[transact.CassetteIndex].Dispense(transact.BanknotesCount)
		}
	}
	return nil
}

// DispenseTransaction TODO: make transaction like if all good
type DispenseTransaction struct {
	CassetteIndex  int
	BanknotesCount int
}
