package utils

import (
	"net/http"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func Pay() {
	api_key := "sk_test_51Hi5f0G3PLuIirSi2m4W9JLbVsHRa0IyaQGPnlvaDINKjZDgMFT1WhHIsjiPJUIXgR0DVlgoT2kqaPh8pox9cvTH00X3P4Sfi7"
	stripe.Key = api_key
	_, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(500),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String("LID"),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String("Nuktu@mail.com")})

	if err != nil {
		print(http.StatusBadRequest, "Payment Unsuccessfull")
		return
	}
}
