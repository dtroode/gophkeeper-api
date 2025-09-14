package model

// Card represents a payment card payload stored in records.
type Card struct {
	CardholderName  string
	Number          string
	Brand           string
	ExpirationMonth int
	ExpirationYear  int
	SecurityCode    string
}
