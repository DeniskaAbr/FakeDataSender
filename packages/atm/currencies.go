package atm

// CurrencyCode Alias hide the real type of the enum
// and users can use it to define the var for accepting enum
type CurrencyCode = int

type Currency struct {
	Rouble CurrencyCode
	USD    CurrencyCode
}

// CurrencyValue Currency for public use
var CurrencyValue = &Currency{
	Rouble: 810,
	USD:    840,
}
