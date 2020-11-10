// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
)

type FraudInfo struct {
	recurlyResponse *ResponseMetadata

	// Kount score
	Score int `json:"score,omitempty"`

	// Kount decision
	Decision string `json:"decision,omitempty"`

	// Kount rules
	RiskRulesTriggered map[string]interface{} `json:"risk_rules_triggered,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *FraudInfo) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *FraudInfo) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type fraudInfoList struct {
	ListMetadata
	Data            []FraudInfo `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *fraudInfoList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *fraudInfoList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// FraudInfoList allows you to paginate FraudInfo objects
type FraudInfoList struct {
	client        HttpCaller
	nextPagePath  string
	genericParams GenericParams

	HasMore bool
	Data    []FraudInfo
}

func NewFraudInfoList(client HttpCaller, nextPagePath string, genericParams GenericParams) *FraudInfoList {
	return &FraudInfoList{
		client:        client,
		nextPagePath:  nextPagePath,
		genericParams: genericParams,
		HasMore:       true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *FraudInfoList) Fetch() error {
	resources := &fraudInfoList{}
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
func (list *FraudInfoList) Count() (*int64, error) {
	resources := &fraudInfoList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
