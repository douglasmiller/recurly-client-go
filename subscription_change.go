// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
	"time"
)

type SubscriptionChange struct {
	recurlyResponse *ResponseMetadata

	// The ID of the Subscription Change.
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The ID of the subscription that is going to be changed.
	SubscriptionId string `json:"subscription_id,omitempty"`

	// Just the important parts.
	Plan PlanMini `json:"plan,omitempty"`

	// These add-ons will be used when the subscription renews.
	AddOns []SubscriptionAddOn `json:"add_ons,omitempty"`

	// Unit amount
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Subscription quantity
	Quantity int `json:"quantity,omitempty"`

	// Subscription shipping details
	Shipping SubscriptionShipping `json:"shipping,omitempty"`

	// Activated at
	ActivateAt time.Time `json:"activate_at,omitempty"`

	// Returns `true` if the subscription change is activated.
	Activated bool `json:"activated,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Invoice Collection
	InvoiceCollection InvoiceCollection `json:"invoice_collection,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *SubscriptionChange) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *SubscriptionChange) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionChangeList struct {
	ListMetadata
	Data            []SubscriptionChange `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionChangeList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionChangeList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionChangeList allows you to paginate SubscriptionChange objects
type SubscriptionChangeList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []SubscriptionChange
}

func NewSubscriptionChangeList(client HttpCaller, nextPagePath string) *SubscriptionChangeList {
	return &SubscriptionChangeList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionChangeList) Fetch() error {
	resources := &subscriptionChangeList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, nil, resources)
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
func (list *SubscriptionChangeList) Count() (*int64, error) {
	resources := &subscriptionChangeList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
