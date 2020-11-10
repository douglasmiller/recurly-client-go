// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
	"time"
)

type CouponRedemption struct {
	recurlyResponse *ResponseMetadata

	// Coupon Redemption ID
	Id string `json:"id,omitempty"`

	// Will always be `coupon`.
	Object string `json:"object,omitempty"`

	// The Account on which the coupon was applied.
	Account AccountMini `json:"account,omitempty"`

	Coupon Coupon `json:"coupon,omitempty"`

	// Coupon Redemption state
	State string `json:"state,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// The amount that was discounted upon the application of the coupon, formatted with the currency.
	Discounted float64 `json:"discounted,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date and time the redemption was removed from the account (un-redeemed).
	RemovedAt time.Time `json:"removed_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponRedemption) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponRedemption) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponRedemptionList struct {
	ListMetadata
	Data            []CouponRedemption `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponRedemptionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponRedemptionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponRedemptionList allows you to paginate CouponRedemption objects
type CouponRedemptionList struct {
	client        HttpCaller
	nextPagePath  string
	genericParams GenericParams

	HasMore bool
	Data    []CouponRedemption
}

func NewCouponRedemptionList(client HttpCaller, nextPagePath string, genericParams GenericParams) *CouponRedemptionList {
	return &CouponRedemptionList{
		client:        client,
		nextPagePath:  nextPagePath,
		genericParams: genericParams,
		HasMore:       true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponRedemptionList) Fetch() error {
	resources := &couponRedemptionList{}
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
func (list *CouponRedemptionList) Count() (*int64, error) {
	resources := &couponRedemptionList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
