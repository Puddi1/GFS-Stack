package handlers

import (
	"fmt"

	"github.com/Puddi1/GFS-Stack/stripe_gfs"
	"github.com/stripe/stripe-go/v74"
	billingportalsession "github.com/stripe/stripe-go/v74/billingportal/session"
	checkoutsession "github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/coupon"
	"github.com/stripe/stripe-go/v74/customer"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
	"github.com/stripe/stripe-go/v74/promotioncode"
	"github.com/stripe/stripe-go/v74/subscription"
	"github.com/stripe/stripe-go/v74/webhookendpoint"
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

// HandleStripeNewCustomer will take a stripe.CustomerParams struct pointer and create a
// new customer, handling any errors and returning the pointer to the customer struct
func HandleStripeNewCustomer(params *stripe.CustomerParams) *stripe.Customer {
	c, err := customer.New(params)
	HandleStripeError(err)
	return c
}

// // Checkout Session // //

// HandleStripeCheckoutSessionCreation will create a new payment session
func HandleStripeCheckoutSessionCreation(params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	s, err := checkoutsession.New(params)
	HandleStripeError(err)
	return s, nil
}

// HandleStripeCheckoutSessionExpire will expire a session based on id
func HandleStripeCheckoutSessionExpire(id string, params *stripe.CheckoutSessionExpireParams) (*stripe.CheckoutSession, error) {
	s, err := checkoutsession.Expire(
		id,
		params,
	)
	HandleStripeError(err)
	return s, nil
}

// HandleStripeGetCheckoutSession will get the checkout session based on id
func HandleStripeGetCheckoutSession(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	s, err := checkoutsession.Get(
		id,
		params,
	)
	HandleStripeError(err)
	return s, nil
}

// HandleStripeGetAllCheckoutSession will get the all checkout session and return them
func HandleStripeGetAllCheckoutSession(
	params *stripe.CheckoutSessionListParams,
	filters []stripe_gfs.FilterStruct,
) *checkoutsession.Iter {
	paramsFiltered := HandleStripeCheckoutSessionListParamsAddFilter(params, filters)

	i := checkoutsession.List(paramsFiltered)

	return i
}

// HandleStripeGetAllCheckoutSessionLineItems gets all the checkout session ine items and return them
func HandleStripeGetAllCheckoutSessionLineItems(
	params *stripe.CheckoutSessionListLineItemsParams,
	filters []stripe_gfs.FilterStruct,
) *checkoutsession.LineItemIter {
	paramsFiltered := HandleStripeCheckoutSessionListLineItemsParamsAddFilter(params, filters)

	i := checkoutsession.ListLineItems(paramsFiltered) // different from stripe api docs

	return i
}

// HandleStripeCheckoutSessionParams will create a new checkout session params struct
func HandleStripeCheckoutSessionParams(
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

// HandleStripeCustomerAddExpand handles all expand params that you want to add to your
// param struct
func HandleStripeCustomerAddExpand(
	params *stripe.CustomerParams,
	expandElements []string,
) (*stripe.CustomerParams, error) {
	for e := range expandElements {
		params.AddExpand(expandElements[e])
	}
	return params, nil
}

// HandleStripeCheckoutSessionListParamsAddFilter handles all filters that you want to add
// to your CheckoutSessionListParams struct
func HandleStripeCheckoutSessionListParamsAddFilter(
	params *stripe.CheckoutSessionListParams,
	filters []stripe_gfs.FilterStruct,
) *stripe.CheckoutSessionListParams {
	for e := range filters {
		params.Filters.AddFilter(
			filters[e].Key,
			filters[e].Op,
			filters[e].Value,
		)
	}
	return params
}

// HandleStripeCheckoutSessionListLineItemsParamsAddFilter handles all filters that you want
// to add to your CheckoutSessionListLineItemsParams struct
func HandleStripeCheckoutSessionListLineItemsParamsAddFilter(
	params *stripe.CheckoutSessionListLineItemsParams,
	filters []stripe_gfs.FilterStruct,
) *stripe.CheckoutSessionListLineItemsParams {
	for e := range filters {
		params.Filters.AddFilter(
			filters[e].Key,
			filters[e].Op,
			filters[e].Value,
		)
	}
	return params
}

// // Customer Portal // //

// HandleStripeCustomerPortalSessionCreation creates a new customer portal session.
func HandleStripeCustomerPortalSessionCreation(id string, returnURL string) (*stripe.BillingPortalSession, error) {
	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(id),
		ReturnURL: stripe.String(returnURL),
	}
	s, err := billingportalsession.New(params)

	HandleStripeError(err)

	return s, nil
}

// // Products // //

// HandleStripeProductCreation will create a new product
func HandleStripeProductCreation(params *stripe.ProductParams) *stripe.Product {
	p, err := product.New(params)
	HandleStripeError(err)
	return p
}

// HandleStripeRetriveProduct will retrive a product object based on object id.
func HandleStripeRetriveProduct(id string) *stripe.Product {
	p, err := product.Get(id, nil)
	HandleStripeError(err)
	return p
}

// HandleStripeProductUpdate will update a product
func HandleStripeProductUpdate(id string, params *stripe.ProductParams, metadata []stripe_gfs.MetadataStruct) *stripe.Product {
	pMetadata := HandleStripeProductMetadata(params, metadata)
	p, err := product.Update(
		id,
		pMetadata,
	)
	HandleStripeError(err)
	return p
}

// HandleStripeRetriveAllProducts will retrive all products object
func HandleStripeRetriveAllProducts(filters []stripe_gfs.FilterStruct) *product.Iter {
	params := &stripe.ProductListParams{}
	pFiltered := HandleStripeProductListParamsAddFilter(params, filters)

	i := product.List(pFiltered)
	return i
}

// HandleStripeProductDeletion will delete a product
func HandleStripeProductDeletion(id string) *stripe.Product {
	p, err := product.Del(id, nil)
	HandleStripeError(err)
	return p
}

// HandleStripeProductSearch will search a product
func HandleStripeProductSearch(id string, params *stripe.ProductSearchParams, query string) *product.SearchIter {
	params.Query = query
	iter := product.Search(params)
	return iter
}

// HandleStripeProductListParamsAddFilter handles all filters that you want to add
// to your ProductListParams struct
func HandleStripeProductListParamsAddFilter(
	params *stripe.ProductListParams,
	filters []stripe_gfs.FilterStruct,
) *stripe.ProductListParams {
	for e := range filters {
		params.Filters.AddFilter(
			filters[e].Key,
			filters[e].Op,
			filters[e].Value,
		)
	}
	return params
}

// HandleStripeProductMetadata handles all metadata fileds that you want to add
// to your ProductParams struct
func HandleStripeProductMetadata(
	params *stripe.ProductParams,
	metadata []stripe_gfs.MetadataStruct,
) *stripe.ProductParams {
	for i := range metadata {
		params.AddMetadata(
			metadata[i].Key,
			metadata[i].Value,
		)
	}
	return params
}

// // Prices // //

// HandleStripeCreatePrice will create a new price of a product based on its id.
func HandleStripeCreatePrice(id string, params *stripe.PriceParams) *stripe.Price {
	p, err := price.New(params)
	HandleStripeError(err)
	return p
}

// HandleStripeRetrivePrice will retrive a price object based on price id.
func HandleStripeRetrivePrice(id string) *stripe.Price {
	p, err := price.Get(id, nil)
	HandleStripeError(err)
	return p
}

// HandleStripeUpdatePrice will update an existing price of a product based on its id.
func HandleStripeUpdatePrice(id string, params *stripe.PriceParams, metadata []stripe_gfs.MetadataStruct) *stripe.Price {
	pMetadata := HandleStripePriceMetadata(params, metadata)
	p, err := price.Update(
		id,
		pMetadata,
	)
	HandleStripeError(err)
	return p
}

// HandleStripeRetriveAllPrices will retrive all prices object
func HandleStripeRetriveAllPrices(filters []stripe_gfs.FilterStruct) *price.Iter {
	params := &stripe.PriceListParams{}
	pFiltered := HandleStripePriceListParamsAddFilter(params, filters)

	i := price.List(pFiltered)
	return i
}

// HandleStripePriceSearch will search a product
func HandleStripePriceSearch(id string, params *stripe.PriceSearchParams, query string) *price.SearchIter {
	params.Query = query
	iter := price.Search(params)
	return iter
}

// HandleStripeProductMetadata handles all metadata fileds that you want to add
// to your ProductParams struct
func HandleStripePriceMetadata(
	params *stripe.PriceParams,
	metadata []stripe_gfs.MetadataStruct,
) *stripe.PriceParams {
	for i := range metadata {
		params.AddMetadata(
			metadata[i].Key,
			metadata[i].Value,
		)
	}
	return params
}

// HandleStripePriceListParamsAddFilter handles all filters that you want to add
// to your PriceListParams struct
func HandleStripePriceListParamsAddFilter(
	params *stripe.PriceListParams,
	filters []stripe_gfs.FilterStruct,
) *stripe.PriceListParams {
	for e := range filters {
		params.Filters.AddFilter(
			filters[e].Key,
			filters[e].Op,
			filters[e].Value,
		)
	}
	return params
}

// // Coupons // //

// HandleStripeCreateCoupon will create a new coupon.
func HandleStripeCreateCoupon(params *stripe.CouponParams) *stripe.Coupon {
	c, err := coupon.New(params)
	HandleStripeError(err)
	return c
}

// HandleStripeRetriveCoupon will retrive a coupon object based on coupon id.
func HandleStripeRetriveCoupon(id string) *stripe.Coupon {
	c, err := coupon.Get(id, nil)
	HandleStripeError(err)
	return c
}

// HandleStripeUpdateCoupon will update a coupon object based on its id.
func HandleStripeUpdateCoupon(id string, params *stripe.CouponParams, metadata []*stripe_gfs.MetadataStruct) *stripe.Coupon {
	pMetadata := HandleStripeCouponMetadata(params, metadata)
	c, err := coupon.Update(id, pMetadata)
	HandleStripeError(err)
	return c
}

// HandleStripeDeleteCoupon will delete a coupon object based on its id.
func HandleStripeDeleteCoupon(id string) *stripe.Coupon {
	c, err := coupon.Del(id, nil)
	HandleStripeError(err)
	return c
}

// HandleStripeRetriveAllCoupons will retrive all coupons object
func HandleStripeRetriveAllCoupons(filters []stripe_gfs.FilterStruct) *coupon.Iter {
	params := &stripe.CouponListParams{}
	pFiltered := HandleStripeCouponListParamsAddFilter(params, filters)

	i := coupon.List(pFiltered)
	return i
}

// HandleStripeCouponListParamsAddFilter handles all filters that you want to add
// to your CouponListParams struct
func HandleStripeCouponListParamsAddFilter(
	params *stripe.CouponListParams,
	filters []stripe_gfs.FilterStruct,
) *stripe.CouponListParams {
	for e := range filters {
		params.Filters.AddFilter(
			filters[e].Key,
			filters[e].Op,
			filters[e].Value,
		)
	}
	return params
}

// HandleStripeCouponMetadata handles all metadata fileds that you want to add
// to your CouponParams struct
func HandleStripeCouponMetadata(
	params *stripe.CouponParams,
	metadata []*stripe_gfs.MetadataStruct,
) *stripe.CouponParams {
	for i := range metadata {
		params.AddMetadata(
			metadata[i].Key,
			metadata[i].Value,
		)
	}
	return params
}

// // Promotions // //

// HandleStripeCreatePromotion will create a new promotion.
func HandleStripeCreatePromotion(params *stripe.PromotionCodeParams) *stripe.PromotionCode {
	pc, err := promotioncode.New(params)
	HandleStripeError(err)
	return pc
}

// HandleStripeRetrivePromotion will retrive a promotion object based on coupon id.
func HandleStripeRetrivePromotion(id string) *stripe.PromotionCode {
	pc, err := promotioncode.Get(id, nil)
	HandleStripeError(err)
	return pc
}

// HandleStripeUpdatePromotion will update a promotion object based on its id.
func HandleStripeUpdatePromotion(id string, params *stripe.PromotionCodeParams, metadata []*stripe_gfs.MetadataStruct) *stripe.PromotionCode {
	pMetadata := HandleStripePromotionMetadata(params, metadata)
	pc, err := promotioncode.Update(
		id,
		pMetadata,
	)
	HandleStripeError(err)
	return pc
}

// HandleStripeRetriveAllPromotions will retrive all promotions object
func HandleStripeRetriveAllPromotions(filters []stripe_gfs.FilterStruct) *promotioncode.Iter {
	params := &stripe.PromotionCodeListParams{}
	pFiltered := HandleStripePromotionListParamsAddFilter(params, filters)
	i := promotioncode.List(pFiltered)
	return i
}

// HandleStripePromotionListParamsAddFilter handles all filters that you want to add
// to your PromotionCodeListParams struct
func HandleStripePromotionListParamsAddFilter(
	params *stripe.PromotionCodeListParams,
	filters []stripe_gfs.FilterStruct,
) *stripe.PromotionCodeListParams {
	for e := range filters {
		params.Filters.AddFilter(
			filters[e].Key,
			filters[e].Op,
			filters[e].Value,
		)
	}
	return params
}

// // HandleStripePromotionMetadata handles all metadata fileds that you want to add
// to your PromotionCodeParams struct
func HandleStripePromotionMetadata(
	params *stripe.PromotionCodeParams,
	metadata []*stripe_gfs.MetadataStruct,
) *stripe.PromotionCodeParams {
	for i := range metadata {
		params.AddMetadata(
			metadata[i].Key,
			metadata[i].Value,
		)
	}
	return params
}

// // Discounts // //

// HandleStripeDeleteCustomerDiscount deletes a customer discount
func HandleStripeDeleteCustomerDiscount(id string) *stripe.Customer {
	cus, err := customer.DeleteDiscount(id, nil)
	HandleStripeError(err)
	return cus
}

// HandleStripeDeleteSubscriptionDiscount deletes a subscription discount
func HandleStripeDeleteSubscriptionDiscount(id string) *stripe.Subscription {
	s, err := subscription.DeleteDiscount(id, nil)
	HandleStripeError(err)
	return s
}

// // Webhooks // //

// HandleStripeCreateWebhook will create a new webhook to listen to a stripe API endpoint
func HandleStripeCreateWebhook(params *stripe.WebhookEndpointParams) *stripe.WebhookEndpoint {
	we, err := webhookendpoint.New(params)
	HandleStripeError(err)
	return we
}

// HandleStripeCreateWebhook will delete a webhook
func HandleStripeRestripeWebhook(id string) *stripe.WebhookEndpoint {
	we, err := webhookendpoint.Del(id, nil)
	HandleStripeError(err)
	return we
}

// HandleStripeCreateWebhook will update a webhook
func HandleStripeUpdateWebhook(id string, params *stripe.WebhookEndpointParams) *stripe.WebhookEndpoint {
	we, err := webhookendpoint.Update(
		id,
		params,
	)
	HandleStripeError(err)
	return we
}

// HandleStripeCreateWebhook will list all webhooks
func HandleStripeListWebhook(filters []stripe_gfs.FilterStruct) *webhookendpoint.Iter {
	params := &stripe.WebhookEndpointListParams{}
	pFiltered := HandleStripeWebhookListParamsAddFilter(params, filters)

	i := webhookendpoint.List(pFiltered)
	return i
}

// HandleStripeCreateWebhook will delete a webhook
func HandleStripeDeleteWebhook(id string) *stripe.WebhookEndpoint {
	we, err := webhookendpoint.Del(
		id,
		nil,
	)
	HandleStripeError(err)
	return we
}

// HandleStripeWebhookListParamsAddFilter handles all filters that you want to add
// to your WebhookEndpointListParams struct
func HandleStripeWebhookListParamsAddFilter(
	params *stripe.WebhookEndpointListParams,
	filters []stripe_gfs.FilterStruct,
) *stripe.WebhookEndpointListParams {
	for e := range filters {
		params.Filters.AddFilter(
			filters[e].Key,
			filters[e].Op,
			filters[e].Value,
		)
	}
	return params
}
