// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
)

type CouponDiscount struct {
	recurlyResponse *ResponseMetadata

	Type string `json:"type,omitempty"`

	// This is only present when `type=percent`.
	Percent int `json:"percent,omitempty"`

	// This is only present when `type=fixed`.
	Currencies []CouponDiscountPricing `json:"currencies,omitempty"`

	// This is only present when `type=free_trial`.
	Trial CouponDiscountTrial `json:"trial,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponDiscount) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponDiscount) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponDiscountList struct {
	ListMetadata
	Data            []CouponDiscount `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponDiscountList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponDiscountList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponDiscountList allows you to paginate CouponDiscount objects
type CouponDiscountList struct {
	client        HttpCaller
	nextPagePath  string
	genericParams GenericParams

	HasMore bool
	Data    []CouponDiscount
}

func NewCouponDiscountList(client HttpCaller, nextPagePath string, genericParams GenericParams) *CouponDiscountList {
	return &CouponDiscountList{
		client:        client,
		nextPagePath:  nextPagePath,
		genericParams: genericParams,
		HasMore:       true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponDiscountList) Fetch() error {
	resources := &couponDiscountList{}
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
func (list *CouponDiscountList) Count() (*int64, error) {
	resources := &couponDiscountList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
