// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
)

type CouponDiscountPricing struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Value of the fixed discount that this coupon applies.
	Amount float64 `json:"amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponDiscountPricing) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponDiscountPricing) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponDiscountPricingList struct {
	ListMetadata
	Data            []CouponDiscountPricing `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponDiscountPricingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponDiscountPricingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponDiscountPricingList allows you to paginate CouponDiscountPricing objects
type CouponDiscountPricingList struct {
	client        HttpCaller
	nextPagePath  string
	genericParams GenericParams

	HasMore bool
	Data    []CouponDiscountPricing
}

func NewCouponDiscountPricingList(client HttpCaller, nextPagePath string, genericParams GenericParams) *CouponDiscountPricingList {
	return &CouponDiscountPricingList{
		client:        client,
		nextPagePath:  nextPagePath,
		genericParams: genericParams,
		HasMore:       true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponDiscountPricingList) Fetch() error {
	resources := &couponDiscountPricingList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, list.genericParams, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Count returns the count of items on the server that match this pager
func (list *CouponDiscountPricingList) Count() (*int64, error) {
	resources := &couponDiscountPricingList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
