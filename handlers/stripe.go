package handlers

import (
	"fmt"
	"net/http"

	"github.com/Puddi1/GFS-Stack/stripe_gfs"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/customer"
)

// Stripe Error Handling // //

// HandleStripeError takes an error returned by a stripe request and it handles it by
// adding Stripe-specific information about what went wrong.
func HandleStripeError(err error) {
	if err != nil {
		// Try to safely cast a generic error to a stripe.Error so that we can get at
		// some additional Stripe-specific information about what went wrong.
		if stripeErr, ok := err.(*stripe.Error); ok {
			// The Code field will contain a basic identifier for the failure.
			switch stripeErr.Code {
			case stripe.ErrorCodeCardDeclined:
			case stripe.ErrorCodeExpiredCard:
			case stripe.ErrorCodeIncorrectCVC:
			case stripe.ErrorCodeIncorrectZip:
				// etc.
			}

			// The Err field can be coerced to a more specific error type with a type
			// assertion. This technique can be used to get more specialized
			// information for certain errors.
			if cardErr, ok := stripeErr.Err.(*stripe.CardError); ok {
				fmt.Printf("Card was declined with code: %v\n", cardErr.DeclineCode)
			} else {
				fmt.Printf("Other Stripe error occurred: %v\n", stripeErr.Error())
			}
		} else {
			fmt.Printf("Other error occurred: %v\n", err.Error())
		}
	}
}

// // Customer Handling // //

// HandleNewCustomer will take a stripe.CustomerParams struct pointer and create a
// new customer, handling any errors and returning the pointer to the customer struct
func HandleNewCustomer(params *stripe.CustomerParams) *stripe.Customer {
	c, err := customer.New(params)
	HandleStripeError(err)
	return c
}

// // Checkout Session // //

// HandleCheckoutSessionCreation will create a new payment session
func HandleCheckoutSessionCreation(params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	s, err := session.New(params)
	HandleStripeError(err)
	return s, nil
}

// HandleCheckoutSessionExpire will expire a session based on id
func HandleCheckoutSessionExpire(id string, params *stripe.CheckoutSessionExpireParams) (*stripe.CheckoutSession, error) {
	s, err := session.Expire(
		id,
		params,
	)
	HandleStripeError(err)
	return s, nil
}

// HandleGetCheckoutSession will get the checkout session based on id
func HandleGetCheckoutSession(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	s, err := session.Get(
		id,
		params,
	)
	HandleStripeError(err)
	return s, nil
}

// HandleGetAllCheckoutSession will get the all checkout session and return them
func HandleGetAllCheckoutSession(
	params *stripe.CheckoutSessionListParams,
	filters []stripe_gfs.FilterStruct,
) (*session.Iter, error) {
	paramsFiltered, err := HandleCheckoutSessionListParamsAddFilter(params, filters)
	_ = err // temporary solution

	i := session.List(paramsFiltered)

	return i, nil
}

// HandleGetAllCheckoutSessionLineItems gets all the checkout session ine items and return them
func HandleGetAllCheckoutSessionLineItems(
	params *stripe.CheckoutSessionListLineItemsParams,
	filters []stripe_gfs.FilterStruct,
) (*session.LineItemIter, error) {
	paramsFiltered, err := HandleCheckoutSessionListLineItemsParamsAddFilter(params, filters)
	_ = err // temporary solution

	i := session.ListLineItems(paramsFiltered) // different from stripe api docs

	return i, err
}

// HandleCheckoutSessionParams will create a new checkout session params struct
func HandleCheckoutSessionParams(
	items []*stripe.CheckoutSessionLineItemParams,
	mode string, cancelURL string, successURL string,
) (*stripe.CheckoutSessionParams, error) {
	params := &stripe.CheckoutSessionParams{
		LineItems:  items,
		Mode:       stripe.String(mode),
		CancelURL:  stripe.String(cancelURL),
		SuccessURL: stripe.String(successURL),
	}
	return params, nil
}

// HandleCustomerAddExpand handles all expand params that you want to add to your
// param struct
func HandleCustomerAddExpand(
	params *stripe.CustomerParams,
	expandElements []string,
) (*stripe.CustomerParams, error) {
	for e := range expandElements {
		params.AddExpand(expandElements[e])
	}
	return params, nil
}

// HandleCheckoutSessionListParamsAddFilter handles all filters that you want to add
// to your CheckoutSessionListParams struct
func HandleCheckoutSessionListParamsAddFilter(
	params *stripe.CheckoutSessionListParams,
	filters []stripe_gfs.FilterStruct,
) (*stripe.CheckoutSessionListParams, error) {
	for e := range filters {
		params.Filters.AddFilter(
			filters[e].Key,
			filters[e].Op,
			filters[e].Value,
		)
	}
	return params, nil
}

// HandleCheckoutSessionListLineItemsParamsAddFilter handles all filters that you want
// to add to your CheckoutSessionListLineItemsParams struct
func HandleCheckoutSessionListLineItemsParamsAddFilter(
	params *stripe.CheckoutSessionListLineItemsParams,
	filters []stripe_gfs.FilterStruct,
) (*stripe.CheckoutSessionListLineItemsParams, error) {
	for e := range filters {
		params.Filters.AddFilter(
			filters[e].Key,
			filters[e].Op,
			filters[e].Value,
		)
	}
	return params, nil
}

// // Customer Portal // //
// HandleCustomerPortalSessionCreation creates a new customer portal session.
// Note: raw http request, stripe-go is giving me problems
func HandleCustomerPortalSessionCreation(id string, returnURL string) (*stripe.CheckoutSession, error) {

	t := "customer=cus_OKTgoI3cQYJLVV return_url=https://example.com/account"
	body := []byte(t)
	res, _ := HandleRequestHTTP(&RequestHTTP{
		MethodHTTP: http.MethodPost,
		Url:        "https://api.stripe.com/v1/billing_portal/sessions",
		Body:       body,
		Headers:    [][2]string{{"Authorization", "Bearer " + stripe.Key}},
	})

	fmt.Print(res)

	var r *stripe.CheckoutSession

	return r, nil
}
