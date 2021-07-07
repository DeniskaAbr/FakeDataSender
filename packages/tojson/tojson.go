package tojson

import (
	"FakeDataSender/packages/atm"
	"encoding/json"
	"strconv"
)

func PackToJSON(a *atm.ATM) string {
	data := new(atmData)
	data.CashIn = append(data.CashIn, struct {
		ValidatorFirmwareTemplate string `json:"validator_firmware_template"`
		Status                    string `json:"status"`
		Accepted                  int    `json:"accepted"`
	}{a.CashIn.ValidatorFirmwareTemplate,
		a.CashIn.Status,
		a.CashIn.Accepted})

	data.MotorCardReader = append(data.MotorCardReader, struct {
		Firmware string `json:"firmware"`
		Status   string `json:"status"`
	}{a.MotorCardReader.Firmware,
		a.MotorCardReader.Status})

	data.ContactlessCardReader = append(data.ContactlessCardReader, struct {
		Firmware string `json:"firmware"`
		Status   string `json:"status"`
	}{a.MotorCardReader.Firmware,
		a.MotorCardReader.Status})

	data.EppPinPad = append(data.EppPinPad, struct {
		SerialKey string `json:"serial_key"`
		Firmware  string `json:"firmware"`
		Status    string `json:"status"`
	}{a.EppPinPad.SerialKey,
		a.EppPinPad.Firmware,
		a.EppPinPad.Status})

	data.SafetyDevice = append(data.SafetyDevice, struct {
		Status   string `json:"status"`
		SafeDoor string `json:"safe_door"`
	}{a.SafetyDevice.Status,
		a.SafetyDevice.SafeDoor})

	mp := make(map[string]Cassettes)

	for i, cassette := range a.CashOut {
		cassetteName := "cassette" + strconv.Itoa(i+1)

		cas1 := Cassette{
			Loaded:       cassette.Loaded,
			Rejected:     cassette.Rejected,
			Dispensed:    cassette.Dispensed,
			Currency:     cassette.Currency,
			Denomination: cassette.Denomination,
			Status:       cassette.Status,
		}
		var mp1 []Cassette
		mp1 = append(mp1, cas1)
		mp[cassetteName] = mp1
	}

	data.CashOut = []CashOuts{}
	data.CashOut = append(data.CashOut, mp)
	marshalledData, _ := json.Marshal(data)
	return string(marshalledData)
}
