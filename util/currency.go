package util

// constants fo all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupportedCurrency checks if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
