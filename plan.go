// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
	"time"
)

type Plan struct {
	recurlyResponse *ResponseMetadata

	// Plan ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Unique code to identify the plan. This is used in Hosted Payment Page URLs and in the invoice exports.
	Code string `json:"code,omitempty"`

	// The current state of the plan.
	State string `json:"state,omitempty"`

	// This name describes your plan and will appear on the Hosted Payment Page and the subscriber's invoice.
	Name string `json:"name,omitempty"`

	// Optional description, not displayed.
	Description string `json:"description,omitempty"`

	// Unit for the plan's billing interval.
	IntervalUnit string `json:"interval_unit,omitempty"`

	// Length of the plan's billing interval in `interval_unit`.
	IntervalLength int `json:"interval_length,omitempty"`

	// Units for the plan's trial period.
	TrialUnit string `json:"trial_unit,omitempty"`

	// Length of plan's trial period in `trial_units`. `0` means `no trial`.
	TrialLength int `json:"trial_length,omitempty"`

	// Allow free trial subscriptions to be created without billing info. Should not be used if billing info is needed for initial invoice due to existing uninvoiced charges or setup fee.
	TrialRequiresBillingInfo bool `json:"trial_requires_billing_info,omitempty"`

	// Automatically terminate subscriptions after a defined number of billing cycles. Number of billing cycles before the plan automatically stops renewing, defaults to `null` for continuous, automatic renewal.
	TotalBillingCycles int `json:"total_billing_cycles,omitempty"`

	// Subscriptions will automatically inherit this value once they are active. If `auto_renew` is `true`, then a subscription will automatically renew its term at renewal. If `auto_renew` is `false`, then a subscription will expire at the end of its term. `auto_renew` can be overridden on the subscription record itself.
	AutoRenew bool `json:"auto_renew,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Setup fee revenue schedule type
	SetupFeeRevenueScheduleType string `json:"setup_fee_revenue_schedule_type,omitempty"`

	// Accounting code for invoice line items for the plan. If no value is provided, it defaults to plan's code.
	AccountingCode string `json:"accounting_code,omitempty"`

	// Accounting code for invoice line items for the plan's setup fee. If no value is provided, it defaults to plan's accounting code.
	SetupFeeAccountingCode string `json:"setup_fee_accounting_code,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the plan is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
	AvalaraTransactionType int `json:"avalara_transaction_type,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the plan is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
	AvalaraServiceType int `json:"avalara_service_type,omitempty"`

	// Used by Avalara, Vertex, and Recurly’s EU VAT tax feature. The tax code values are specific to each tax system. If you are using Recurly’s EU VAT feature you can use `unknown`, `physical`, or `digital`.
	TaxCode string `json:"tax_code,omitempty"`

	// `true` exempts tax on the plan, `false` applies tax on the plan.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// Pricing
	Currencies []PlanPricing `json:"currencies,omitempty"`

	// Hosted pages settings
	HostedPages PlanHostedPages `json:"hosted_pages,omitempty"`

	// Used to determine whether items can be assigned as add-ons to individual subscriptions.
	// If `true`, items can be assigned as add-ons to individual subscription add-ons.
	// If `false`, only plan add-ons can be used.
	AllowAnyItemOnSubscriptions bool `json:"allow_any_item_on_subscriptions,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Plan) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Plan) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planList struct {
	ListMetadata
	Data            []Plan `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanList allows you to paginate Plan objects
type PlanList struct {
	client        HttpCaller
	nextPagePath  string
	genericParams GenericParams

	HasMore bool
	Data    []Plan
}

func NewPlanList(client HttpCaller, nextPagePath string, genericParams GenericParams) *PlanList {
	return &PlanList{
		client:        client,
		nextPagePath:  nextPagePath,
		genericParams: genericParams,
		HasMore:       true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanList) Fetch() error {
	resources := &planList{}
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
func (list *PlanList) Count() (*int64, error) {
	resources := &planList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
