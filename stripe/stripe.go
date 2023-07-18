package stripe

import (
	"github.com/Puddi1/GFS-Stack/env"
	"github.com/stripe/stripe-go/v74"
)

func Init_stripe () {
	stripe.Key = env.ENVs["STRIPE_API_KEY"]
}