package stripe_gfs

import (
	"github.com/Puddi1/GFS-Stack/env"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/client"
)

const (
	PAYMENT      = "payment"
	SETUP        = "setup"
	SUBSCRIPTION = "subscription"
)

type FilterStruct struct {
	Key   string `json:"key"`
	Op    string `json:"op"`
	Value string `json:"value"`
}

type MetadataStruct struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var SC *client.API

func Init_stripe() {
	stripe.Key = env.ENVs["STRIPE_API_PRIVATE_KEY"]
	SC = &client.API{}
	SC.Init(stripe.Key, nil)
}
