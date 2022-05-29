package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	JPY = "JPY"
	CNY = "CNY"
)

var supportedCurrency = [...]string{USD, EUR, CAD, JPY, CNY}

func IsSupportedCurrency(currency string) bool {
	for _, c := range supportedCurrency {
		if c == currency {
			return true
		}
	}
	return false
}
