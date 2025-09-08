package util

const (
	USD = "USD"
	INR = "INR"
)

// IsSupportedCurrency returns true if currency is supported:
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, INR:
		return true
	}
	return false
}
