package tojson

// Cassette TODO: remove or use
type Cassette struct {
	Loaded       int    `json:"loaded"`
	Rejected     int    `json:"rejected"`
	Dispensed    int    `json:"dispensed"`
	Currency     int    `json:"currency"`
	Denomination int    `json:"denomination"`
	Status       string `json:"status"`
}

type Cassettes []Cassette

type CashOuts map[string]Cassettes

type atmData struct {
	CashOut []CashOuts `json:"cash_out"`
	CashIn  []struct {
		ValidatorFirmwareTemplate string `json:"validator_firmware_template"`
		Status                    string `json:"status"`
		Accepted                  int    `json:"accepted"`
	} `json:"cash_in"`
	MotorCardReader []struct {
		Firmware string `json:"firmware"`
		Status   string `json:"status"`
	} `json:"motor_CardReader"`
	ContactlessCardReader []struct {
		Firmware string `json:"firmware"`
		Status   string `json:"status"`
	} `json:"contactless_CardReader"`
	EppPinPad []struct {
		SerialKey string `json:"serial_key"`
		Firmware  string `json:"firmware"`
		Status    string `json:"status"`
	} `json:"epp_PinPad"`
	SafetyDevice []struct {
		Status   string `json:"status"`
		SafeDoor string `json:"safe_door"`
	} `json:"safety_device"`
}
