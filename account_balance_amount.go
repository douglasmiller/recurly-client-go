// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
)

type AccountBalanceAmount struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Total amount the account is past due.
	Amount float64 `json:"amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountBalanceAmount) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountBalanceAmount) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountBalanceAmountList struct {
	ListMetadata
	Data            []AccountBalanceAmount `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountBalanceAmountList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountBalanceAmountList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountBalanceAmountList allows you to paginate AccountBalanceAmount objects
type AccountBalanceAmountList struct {
	client        HttpCaller
	nextPagePath  string
	genericParams GenericParams

	HasMore bool
	Data    []AccountBalanceAmount
}

func NewAccountBalanceAmountList(client HttpCaller, nextPagePath string, genericParams GenericParams) *AccountBalanceAmountList {
	return &AccountBalanceAmountList{
		client:        client,
		nextPagePath:  nextPagePath,
		genericParams: genericParams,
		HasMore:       true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountBalanceAmountList) Fetch() error {
	resources := &accountBalanceAmountList{}
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
func (list *AccountBalanceAmountList) Count() (*int64, error) {
	resources := &accountBalanceAmountList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
