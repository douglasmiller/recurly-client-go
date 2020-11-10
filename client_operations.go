// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// APIVersion is the current Recurly API Version
	APIVersion = "v2020-01-01"
)

type ClientInterface interface {
	ListSites(params *ListSitesParams) (*SiteList, error)

	GetSite(siteId string, params *GetSiteParams) (*Site, error)

	ListAccounts(params *ListAccountsParams) (*AccountList, error)

	CreateAccount(body *AccountCreate, params *CreateAccountParams) (*Account, error)

	GetAccount(accountId string, params *GetAccountParams) (*Account, error)

	UpdateAccount(accountId string, body *AccountUpdate, params *UpdateAccountParams) (*Account, error)

	DeactivateAccount(accountId string, params *DeactivateAccountParams) (*Account, error)

	GetAccountAcquisition(accountId string, params *GetAccountAcquisitionParams) (*AccountAcquisition, error)

	UpdateAccountAcquisition(accountId string, body *AccountAcquisitionUpdatable, params *UpdateAccountAcquisitionParams) (*AccountAcquisition, error)

	RemoveAccountAcquisition(accountId string, params *RemoveAccountAcquisitionParams) (*Empty, error)

	ReactivateAccount(accountId string, params *ReactivateAccountParams) (*Account, error)

	GetAccountBalance(accountId string, params *GetAccountBalanceParams) (*AccountBalance, error)

	GetBillingInfo(accountId string, params *GetBillingInfoParams) (*BillingInfo, error)

	UpdateBillingInfo(accountId string, body *BillingInfoCreate, params *UpdateBillingInfoParams) (*BillingInfo, error)

	RemoveBillingInfo(accountId string, params *RemoveBillingInfoParams) (*Empty, error)

	ListBillingInfos(accountId string, params *ListBillingInfosParams) (*BillingInfoList, error)

	CreateBillingInfo(accountId string, body *BillingInfoCreate, params *CreateBillingInfoParams) (*BillingInfo, error)

	GetABillingInfo(accountId string, billingInfoId string, params *GetABillingInfoParams) (*BillingInfo, error)

	UpdateABillingInfo(accountId string, billingInfoId string, body *BillingInfoCreate, params *UpdateABillingInfoParams) (*BillingInfo, error)

	RemoveABillingInfo(accountId string, billingInfoId string, params *RemoveABillingInfoParams) (*Empty, error)

	ListAccountCouponRedemptions(accountId string, params *ListAccountCouponRedemptionsParams) (*CouponRedemptionList, error)

	ListActiveCouponRedemptions(accountId string, params *ListActiveCouponRedemptionsParams) (*CouponRedemptionList, error)

	CreateCouponRedemption(accountId string, body *CouponRedemptionCreate, params *CreateCouponRedemptionParams) (*CouponRedemption, error)

	RemoveCouponRedemption(accountId string, params *RemoveCouponRedemptionParams) (*CouponRedemption, error)

	ListAccountCreditPayments(accountId string, params *ListAccountCreditPaymentsParams) (*CreditPaymentList, error)

	ListAccountInvoices(accountId string, params *ListAccountInvoicesParams) (*InvoiceList, error)

	CreateInvoice(accountId string, body *InvoiceCreate, params *CreateInvoiceParams) (*InvoiceCollection, error)

	PreviewInvoice(accountId string, body *InvoiceCreate, params *PreviewInvoiceParams) (*InvoiceCollection, error)

	ListAccountLineItems(accountId string, params *ListAccountLineItemsParams) (*LineItemList, error)

	CreateLineItem(accountId string, body *LineItemCreate, params *CreateLineItemParams) (*LineItem, error)

	ListAccountNotes(accountId string, params *ListAccountNotesParams) (*AccountNoteList, error)

	GetAccountNote(accountId string, accountNoteId string, params *GetAccountNoteParams) (*AccountNote, error)

	ListShippingAddresses(accountId string, params *ListShippingAddressesParams) (*ShippingAddressList, error)

	CreateShippingAddress(accountId string, body *ShippingAddressCreate, params *CreateShippingAddressParams) (*ShippingAddress, error)

	GetShippingAddress(accountId string, shippingAddressId string, params *GetShippingAddressParams) (*ShippingAddress, error)

	UpdateShippingAddress(accountId string, shippingAddressId string, body *ShippingAddressUpdate, params *UpdateShippingAddressParams) (*ShippingAddress, error)

	RemoveShippingAddress(accountId string, shippingAddressId string, params *RemoveShippingAddressParams) (*Empty, error)

	ListAccountSubscriptions(accountId string, params *ListAccountSubscriptionsParams) (*SubscriptionList, error)

	ListAccountTransactions(accountId string, params *ListAccountTransactionsParams) (*TransactionList, error)

	ListChildAccounts(accountId string, params *ListChildAccountsParams) (*AccountList, error)

	ListAccountAcquisition(params *ListAccountAcquisitionParams) (*AccountAcquisitionList, error)

	ListCoupons(params *ListCouponsParams) (*CouponList, error)

	CreateCoupon(body *CouponCreate, params *CreateCouponParams) (*Coupon, error)

	GetCoupon(couponId string, params *GetCouponParams) (*Coupon, error)

	UpdateCoupon(couponId string, body *CouponUpdate, params *UpdateCouponParams) (*Coupon, error)

	DeactivateCoupon(couponId string, params *DeactivateCouponParams) (*Coupon, error)

	RestoreCoupon(couponId string, body *CouponUpdate, params *RestoreCouponParams) (*Coupon, error)

	ListUniqueCouponCodes(couponId string, params *ListUniqueCouponCodesParams) (*UniqueCouponCodeList, error)

	ListCreditPayments(params *ListCreditPaymentsParams) (*CreditPaymentList, error)

	GetCreditPayment(creditPaymentId string, params *GetCreditPaymentParams) (*CreditPayment, error)

	ListCustomFieldDefinitions(params *ListCustomFieldDefinitionsParams) (*CustomFieldDefinitionList, error)

	GetCustomFieldDefinition(customFieldDefinitionId string, params *GetCustomFieldDefinitionParams) (*CustomFieldDefinition, error)

	ListItems(params *ListItemsParams) (*ItemList, error)

	CreateItem(body *ItemCreate, params *CreateItemParams) (*Item, error)

	GetItem(itemId string, params *GetItemParams) (*Item, error)

	UpdateItem(itemId string, body *ItemUpdate, params *UpdateItemParams) (*Item, error)

	DeactivateItem(itemId string, params *DeactivateItemParams) (*Item, error)

	ReactivateItem(itemId string, params *ReactivateItemParams) (*Item, error)

	ListMeasuredUnit(params *ListMeasuredUnitParams) (*MeasuredUnitList, error)

	CreateMeasuredUnit(body *MeasuredUnitCreate, params *CreateMeasuredUnitParams) (*MeasuredUnit, error)

	GetMeasuredUnit(measuredUnitId string, params *GetMeasuredUnitParams) (*MeasuredUnit, error)

	UpdateMeasuredUnit(measuredUnitId string, body *MeasuredUnitUpdate, params *UpdateMeasuredUnitParams) (*MeasuredUnit, error)

	RemoveMeasuredUnit(measuredUnitId string, params *RemoveMeasuredUnitParams) (*MeasuredUnit, error)

	ListInvoices(params *ListInvoicesParams) (*InvoiceList, error)

	GetInvoice(invoiceId string, params *GetInvoiceParams) (*Invoice, error)

	PutInvoice(invoiceId string, body *InvoiceUpdatable, params *PutInvoiceParams) (*Invoice, error)

	CollectInvoice(invoiceId string, params *CollectInvoiceParams) (*Invoice, error)

	FailInvoice(invoiceId string, params *FailInvoiceParams) (*Invoice, error)

	MarkInvoiceSuccessful(invoiceId string, params *MarkInvoiceSuccessfulParams) (*Invoice, error)

	ReopenInvoice(invoiceId string, params *ReopenInvoiceParams) (*Invoice, error)

	VoidInvoice(invoiceId string, params *VoidInvoiceParams) (*Invoice, error)

	RecordExternalTransaction(invoiceId string, body *ExternalTransaction, params *RecordExternalTransactionParams) (*Transaction, error)

	ListInvoiceLineItems(invoiceId string, params *ListInvoiceLineItemsParams) (*LineItemList, error)

	ListInvoiceCouponRedemptions(invoiceId string, params *ListInvoiceCouponRedemptionsParams) (*CouponRedemptionList, error)

	ListRelatedInvoices(invoiceId string, params *ListRelatedInvoicesParams) (*InvoiceList, error)

	RefundInvoice(invoiceId string, body *InvoiceRefund, params *RefundInvoiceParams) (*Invoice, error)

	ListLineItems(params *ListLineItemsParams) (*LineItemList, error)

	GetLineItem(lineItemId string, params *GetLineItemParams) (*LineItem, error)

	RemoveLineItem(lineItemId string, params *RemoveLineItemParams) (*Empty, error)

	ListPlans(params *ListPlansParams) (*PlanList, error)

	CreatePlan(body *PlanCreate, params *CreatePlanParams) (*Plan, error)

	GetPlan(planId string, params *GetPlanParams) (*Plan, error)

	UpdatePlan(planId string, body *PlanUpdate, params *UpdatePlanParams) (*Plan, error)

	RemovePlan(planId string, params *RemovePlanParams) (*Plan, error)

	ListPlanAddOns(planId string, params *ListPlanAddOnsParams) (*AddOnList, error)

	CreatePlanAddOn(planId string, body *AddOnCreate, params *CreatePlanAddOnParams) (*AddOn, error)

	GetPlanAddOn(planId string, addOnId string, params *GetPlanAddOnParams) (*AddOn, error)

	UpdatePlanAddOn(planId string, addOnId string, body *AddOnUpdate, params *UpdatePlanAddOnParams) (*AddOn, error)

	RemovePlanAddOn(planId string, addOnId string, params *RemovePlanAddOnParams) (*AddOn, error)

	ListAddOns(params *ListAddOnsParams) (*AddOnList, error)

	GetAddOn(addOnId string, params *GetAddOnParams) (*AddOn, error)

	ListShippingMethods(params *ListShippingMethodsParams) (*ShippingMethodList, error)

	CreateShippingMethod(body *ShippingMethodCreate, params *CreateShippingMethodParams) (*ShippingMethod, error)

	GetShippingMethod(shippingMethodId string, params *GetShippingMethodParams) (*ShippingMethod, error)

	UpdateShippingMethod(shippingMethodId string, body *ShippingMethodUpdate, params *UpdateShippingMethodParams) (*ShippingMethod, error)

	DeactivateShippingMethod(shippingMethodId string, params *DeactivateShippingMethodParams) (*ShippingMethod, error)

	ListSubscriptions(params *ListSubscriptionsParams) (*SubscriptionList, error)

	CreateSubscription(body *SubscriptionCreate, params *CreateSubscriptionParams) (*Subscription, error)

	GetSubscription(subscriptionId string, params *GetSubscriptionParams) (*Subscription, error)

	ModifySubscription(subscriptionId string, body *SubscriptionUpdate, params *ModifySubscriptionParams) (*Subscription, error)

	TerminateSubscription(subscriptionId string, params *TerminateSubscriptionParams) (*Subscription, error)

	CancelSubscription(subscriptionId string, params *CancelSubscriptionParams) (*Subscription, error)

	ReactivateSubscription(subscriptionId string, params *ReactivateSubscriptionParams) (*Subscription, error)

	PauseSubscription(subscriptionId string, body *SubscriptionPause, params *PauseSubscriptionParams) (*Subscription, error)

	ResumeSubscription(subscriptionId string, params *ResumeSubscriptionParams) (*Subscription, error)

	ConvertTrial(subscriptionId string, params *ConvertTrialParams) (*Subscription, error)

	GetSubscriptionChange(subscriptionId string, params *GetSubscriptionChangeParams) (*SubscriptionChange, error)

	CreateSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate, params *CreateSubscriptionChangeParams) (*SubscriptionChange, error)

	RemoveSubscriptionChange(subscriptionId string, params *RemoveSubscriptionChangeParams) (*Empty, error)

	PreviewSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate, params *PreviewSubscriptionChangeParams) (*SubscriptionChange, error)

	ListSubscriptionInvoices(subscriptionId string, params *ListSubscriptionInvoicesParams) (*InvoiceList, error)

	ListSubscriptionLineItems(subscriptionId string, params *ListSubscriptionLineItemsParams) (*LineItemList, error)

	ListSubscriptionCouponRedemptions(subscriptionId string, params *ListSubscriptionCouponRedemptionsParams) (*CouponRedemptionList, error)

	ListUsage(subscriptionId string, addOnId string, params *ListUsageParams) (*UsageList, error)

	CreateUsage(subscriptionId string, addOnId string, body *UsageCreate, params *CreateUsageParams) (*Usage, error)

	GetUsage(usageId string, params *GetUsageParams) (*Usage, error)

	UpdateUsage(usageId string, body *UsageCreate, params *UpdateUsageParams) (*Usage, error)

	RemoveUsage(usageId string, params *RemoveUsageParams) (*Empty, error)

	ListTransactions(params *ListTransactionsParams) (*TransactionList, error)

	GetTransaction(transactionId string, params *GetTransactionParams) (*Transaction, error)

	GetUniqueCouponCode(uniqueCouponCodeId string, params *GetUniqueCouponCodeParams) (*UniqueCouponCode, error)

	DeactivateUniqueCouponCode(uniqueCouponCodeId string, params *DeactivateUniqueCouponCodeParams) (*UniqueCouponCode, error)

	ReactivateUniqueCouponCode(uniqueCouponCodeId string, params *ReactivateUniqueCouponCodeParams) (*UniqueCouponCode, error)

	CreatePurchase(body *PurchaseCreate, params *CreatePurchaseParams) (*InvoiceCollection, error)

	PreviewPurchase(body *PurchaseCreate, params *PreviewPurchaseParams) (*InvoiceCollection, error)

	GetExportDates(params *GetExportDatesParams) (*ExportDates, error)

	GetExportFiles(exportDate string, params *GetExportFilesParams) (*ExportFiles, error)
}

type ListSitesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// State - Filter by state.
	State *string
}

func (list *ListSitesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListSitesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListSites List sites
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_sites
//
// Returns: A list of sites.
func (c *Client) ListSites(params *ListSitesParams) (*SiteList, error) {
	path, err := c.InterpolatePath("/sites")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewSiteList(c, path, params), nil
}

type GetSiteParams struct {
	Params
}

func (list *GetSiteParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetSite Fetch a site
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_site
//
// Returns: A site.
func (c *Client) GetSite(siteId string, params *GetSiteParams) (*Site, error) {
	path, err := c.InterpolatePath("/sites/{site_id}", siteId)
	if err != nil {
		return nil, err
	}
	result := &Site{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Email - Filter for accounts with this exact email address. A blank value will return accounts with both `null` and `""` email addresses. Note that multiple accounts can share one email address.
	Email *string

	// Subscriber - Filter for accounts with or without a subscription in the `active`,
	// `canceled`, or `future` state.
	Subscriber *bool

	// PastDue - Filter for accounts with an invoice in the `past_due` state.
	PastDue *string
}

func (list *ListAccountsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListAccountsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Email != nil {
		options = append(options, KeyValue{Key: "email", Value: *list.Email})
	}

	if list.Subscriber != nil {
		options = append(options, KeyValue{Key: "subscriber", Value: strconv.FormatBool(*list.Subscriber)})
	}

	if list.PastDue != nil {
		options = append(options, KeyValue{Key: "past_due", Value: *list.PastDue})
	}

	return options
}

// ListAccounts List a site's accounts
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_accounts
//
// Returns: A list of the site's accounts.
func (c *Client) ListAccounts(params *ListAccountsParams) (*AccountList, error) {
	path, err := c.InterpolatePath("/accounts")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewAccountList(c, path, params), nil
}

type CreateAccountParams struct {
	Params
}

func (list *CreateAccountParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateAccount Create an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_account
//
// Returns: An account.
func (c *Client) CreateAccount(body *AccountCreate, params *CreateAccountParams) (*Account, error) {
	path, err := c.InterpolatePath("/accounts")
	if err != nil {
		return nil, err
	}
	result := &Account{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetAccountParams struct {
	Params
}

func (list *GetAccountParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetAccount Fetch an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_account
//
// Returns: An account.
func (c *Client) GetAccount(accountId string, params *GetAccountParams) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}", accountId)
	if err != nil {
		return nil, err
	}
	result := &Account{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdateAccountParams struct {
	Params
}

func (list *UpdateAccountParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdateAccount Modify an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_account
//
// Returns: An account.
func (c *Client) UpdateAccount(accountId string, body *AccountUpdate, params *UpdateAccountParams) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}", accountId)
	if err != nil {
		return nil, err
	}
	result := &Account{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type DeactivateAccountParams struct {
	Params
}

func (list *DeactivateAccountParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// DeactivateAccount Deactivate an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/deactivate_account
//
// Returns: An account.
func (c *Client) DeactivateAccount(accountId string, params *DeactivateAccountParams) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}", accountId)
	if err != nil {
		return nil, err
	}
	result := &Account{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetAccountAcquisitionParams struct {
	Params
}

func (list *GetAccountAcquisitionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetAccountAcquisition Fetch an account's acquisition data
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_account_acquisition
//
// Returns: An account's acquisition data.
func (c *Client) GetAccountAcquisition(accountId string, params *GetAccountAcquisitionParams) (*AccountAcquisition, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	if err != nil {
		return nil, err
	}
	result := &AccountAcquisition{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdateAccountAcquisitionParams struct {
	Params
}

func (list *UpdateAccountAcquisitionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdateAccountAcquisition Update an account's acquisition data
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_account_acquisition
//
// Returns: An account's updated acquisition data.
func (c *Client) UpdateAccountAcquisition(accountId string, body *AccountAcquisitionUpdatable, params *UpdateAccountAcquisitionParams) (*AccountAcquisition, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	if err != nil {
		return nil, err
	}
	result := &AccountAcquisition{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RemoveAccountAcquisitionParams struct {
	Params
}

func (list *RemoveAccountAcquisitionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RemoveAccountAcquisition Remove an account's acquisition data
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_account_acquisition
//
// Returns: Acquisition data was succesfully deleted.
func (c *Client) RemoveAccountAcquisition(accountId string, params *RemoveAccountAcquisitionParams) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ReactivateAccountParams struct {
	Params
}

func (list *ReactivateAccountParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// ReactivateAccount Reactivate an inactive account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/reactivate_account
//
// Returns: An account.
func (c *Client) ReactivateAccount(accountId string, params *ReactivateAccountParams) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/reactivate", accountId)
	if err != nil {
		return nil, err
	}
	result := &Account{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetAccountBalanceParams struct {
	Params
}

func (list *GetAccountBalanceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetAccountBalance Fetch an account's balance and past due status
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_account_balance
//
// Returns: An account's balance.
func (c *Client) GetAccountBalance(accountId string, params *GetAccountBalanceParams) (*AccountBalance, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/balance", accountId)
	if err != nil {
		return nil, err
	}
	result := &AccountBalance{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetBillingInfoParams struct {
	Params
}

func (list *GetBillingInfoParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetBillingInfo Fetch an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_billing_info
//
// Returns: An account's billing information.
func (c *Client) GetBillingInfo(accountId string, params *GetBillingInfoParams) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	if err != nil {
		return nil, err
	}
	result := &BillingInfo{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdateBillingInfoParams struct {
	Params
}

func (list *UpdateBillingInfoParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdateBillingInfo Set an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_billing_info
//
// Returns: Updated billing information.
func (c *Client) UpdateBillingInfo(accountId string, body *BillingInfoCreate, params *UpdateBillingInfoParams) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	if err != nil {
		return nil, err
	}
	result := &BillingInfo{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RemoveBillingInfoParams struct {
	Params
}

func (list *RemoveBillingInfoParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RemoveBillingInfo Remove an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_billing_info
//
// Returns: Billing information deleted
func (c *Client) RemoveBillingInfo(accountId string, params *RemoveBillingInfoParams) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListBillingInfosParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListBillingInfosParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListBillingInfosParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListBillingInfos Get the list of billing information associated with an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_billing_infos
//
// Returns: A list of the the billing information for an account's
func (c *Client) ListBillingInfos(accountId string, params *ListBillingInfosParams) (*BillingInfoList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewBillingInfoList(c, path, params), nil
}

type CreateBillingInfoParams struct {
	Params
}

func (list *CreateBillingInfoParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateBillingInfo Set an account's billing information when the wallet feature is enabled
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_billing_info
//
// Returns: Updated billing information.
func (c *Client) CreateBillingInfo(accountId string, body *BillingInfoCreate, params *CreateBillingInfoParams) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos", accountId)
	if err != nil {
		return nil, err
	}
	result := &BillingInfo{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetABillingInfoParams struct {
	Params
}

func (list *GetABillingInfoParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetABillingInfo Fetch a billing info
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_a_billing_info
//
// Returns: A billing info.
func (c *Client) GetABillingInfo(accountId string, billingInfoId string, params *GetABillingInfoParams) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos/{billing_info_id}", accountId, billingInfoId)
	if err != nil {
		return nil, err
	}
	result := &BillingInfo{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdateABillingInfoParams struct {
	Params
}

func (list *UpdateABillingInfoParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdateABillingInfo Update an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_a_billing_info
//
// Returns: Updated billing information.
func (c *Client) UpdateABillingInfo(accountId string, billingInfoId string, body *BillingInfoCreate, params *UpdateABillingInfoParams) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos/{billing_info_id}", accountId, billingInfoId)
	if err != nil {
		return nil, err
	}
	result := &BillingInfo{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RemoveABillingInfoParams struct {
	Params
}

func (list *RemoveABillingInfoParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RemoveABillingInfo Remove an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_a_billing_info
//
// Returns: Billing information deleted
func (c *Client) RemoveABillingInfo(accountId string, billingInfoId string, params *RemoveABillingInfoParams) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos/{billing_info_id}", accountId, billingInfoId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountCouponRedemptionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListAccountCouponRedemptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListAccountCouponRedemptionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListAccountCouponRedemptions Show the coupon redemptions for an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_coupon_redemptions
//
// Returns: A list of the the coupon redemptions on an account.
func (c *Client) ListAccountCouponRedemptions(accountId string, params *ListAccountCouponRedemptionsParams) (*CouponRedemptionList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewCouponRedemptionList(c, path, params), nil
}

type ListActiveCouponRedemptionsParams struct {
	Params
}

func (list *ListActiveCouponRedemptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// ListActiveCouponRedemptions Show the coupon redemptions that are active on an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_active_coupon_redemptions
//
// Returns: Active coupon redemptions on an account.
func (c *Client) ListActiveCouponRedemptions(accountId string, params *ListActiveCouponRedemptionsParams) (*CouponRedemptionList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	if err != nil {
		return nil, err
	}
	params.IsListOp = true
	return NewCouponRedemptionList(c, path, params), nil
}

type CreateCouponRedemptionParams struct {
	Params
}

func (list *CreateCouponRedemptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateCouponRedemption Generate an active coupon redemption on an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_coupon_redemption
//
// Returns: Returns the new coupon redemption.
func (c *Client) CreateCouponRedemption(accountId string, body *CouponRedemptionCreate, params *CreateCouponRedemptionParams) (*CouponRedemption, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	if err != nil {
		return nil, err
	}
	result := &CouponRedemption{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RemoveCouponRedemptionParams struct {
	Params
}

func (list *RemoveCouponRedemptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RemoveCouponRedemption Delete the active coupon redemption from an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_coupon_redemption
//
// Returns: Coupon redemption deleted.
func (c *Client) RemoveCouponRedemption(accountId string, params *RemoveCouponRedemptionParams) (*CouponRedemption, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	if err != nil {
		return nil, err
	}
	result := &CouponRedemption{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountCreditPaymentsParams struct {
	Params

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListAccountCreditPaymentsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListAccountCreditPaymentsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListAccountCreditPayments List an account's credit payments
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_credit_payments
//
// Returns: A list of the account's credit payments.
func (c *Client) ListAccountCreditPayments(accountId string, params *ListAccountCreditPaymentsParams) (*CreditPaymentList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/credit_payments", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewCreditPaymentList(c, path, params), nil
}

type ListAccountInvoicesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Type - Filter by type when:
	// - `type=charge`, only charge invoices will be returned.
	// - `type=credit`, only credit invoices will be returned.
	// - `type=non-legacy`, only charge and credit invoices will be returned.
	// - `type=legacy`, only legacy invoices will be returned.
	Type *string
}

func (list *ListAccountInvoicesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListAccountInvoicesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListAccountInvoices List an account's invoices
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_invoices
//
// Returns: A list of the account's invoices.
func (c *Client) ListAccountInvoices(accountId string, params *ListAccountInvoicesParams) (*InvoiceList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/invoices", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewInvoiceList(c, path, params), nil
}

type CreateInvoiceParams struct {
	Params
}

func (list *CreateInvoiceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateInvoice Create an invoice for pending line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_invoice
//
// Returns: Returns the new invoices.
func (c *Client) CreateInvoice(accountId string, body *InvoiceCreate, params *CreateInvoiceParams) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/invoices", accountId)
	if err != nil {
		return nil, err
	}
	result := &InvoiceCollection{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type PreviewInvoiceParams struct {
	Params
}

func (list *PreviewInvoiceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// PreviewInvoice Preview new invoice for pending line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/preview_invoice
//
// Returns: Returns the invoice previews.
func (c *Client) PreviewInvoice(accountId string, body *InvoiceCreate, params *PreviewInvoiceParams) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/invoices/preview", accountId)
	if err != nil {
		return nil, err
	}
	result := &InvoiceCollection{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountLineItemsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Original - Filter by original field.
	Original *string

	// State - Filter by state field.
	State *string

	// Type - Filter by type field.
	Type *string
}

func (list *ListAccountLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListAccountLineItemsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Original != nil {
		options = append(options, KeyValue{Key: "original", Value: *list.Original})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListAccountLineItems List an account's line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_line_items
//
// Returns: A list of the account's line items.
func (c *Client) ListAccountLineItems(accountId string, params *ListAccountLineItemsParams) (*LineItemList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/line_items", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewLineItemList(c, path, params), nil
}

type CreateLineItemParams struct {
	Params
}

func (list *CreateLineItemParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateLineItem Create a new line item for the account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_line_item
//
// Returns: Returns the new line item.
func (c *Client) CreateLineItem(accountId string, body *LineItemCreate, params *CreateLineItemParams) (*LineItem, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/line_items", accountId)
	if err != nil {
		return nil, err
	}
	result := &LineItem{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountNotesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string
}

func (list *ListAccountNotesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListAccountNotesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	return options
}

// ListAccountNotes Fetch a list of an account's notes
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_notes
//
// Returns: A list of an account's notes.
func (c *Client) ListAccountNotes(accountId string, params *ListAccountNotesParams) (*AccountNoteList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/notes", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewAccountNoteList(c, path, params), nil
}

type GetAccountNoteParams struct {
	Params
}

func (list *GetAccountNoteParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetAccountNote Fetch an account note
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_account_note
//
// Returns: An account note.
func (c *Client) GetAccountNote(accountId string, accountNoteId string, params *GetAccountNoteParams) (*AccountNote, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/notes/{account_note_id}", accountId, accountNoteId)
	if err != nil {
		return nil, err
	}
	result := &AccountNote{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListShippingAddressesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListShippingAddressesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListShippingAddressesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListShippingAddresses Fetch a list of an account's shipping addresses
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_shipping_addresses
//
// Returns: A list of an account's shipping addresses.
func (c *Client) ListShippingAddresses(accountId string, params *ListShippingAddressesParams) (*ShippingAddressList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewShippingAddressList(c, path, params), nil
}

type CreateShippingAddressParams struct {
	Params
}

func (list *CreateShippingAddressParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateShippingAddress Create a new shipping address for the account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_shipping_address
//
// Returns: Returns the new shipping address.
func (c *Client) CreateShippingAddress(accountId string, body *ShippingAddressCreate, params *CreateShippingAddressParams) (*ShippingAddress, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses", accountId)
	if err != nil {
		return nil, err
	}
	result := &ShippingAddress{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetShippingAddressParams struct {
	Params
}

func (list *GetShippingAddressParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetShippingAddress Fetch an account's shipping address
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_shipping_address
//
// Returns: A shipping address.
func (c *Client) GetShippingAddress(accountId string, shippingAddressId string, params *GetShippingAddressParams) (*ShippingAddress, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	if err != nil {
		return nil, err
	}
	result := &ShippingAddress{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdateShippingAddressParams struct {
	Params
}

func (list *UpdateShippingAddressParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdateShippingAddress Update an account's shipping address
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_shipping_address
//
// Returns: The updated shipping address.
func (c *Client) UpdateShippingAddress(accountId string, shippingAddressId string, body *ShippingAddressUpdate, params *UpdateShippingAddressParams) (*ShippingAddress, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	if err != nil {
		return nil, err
	}
	result := &ShippingAddress{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RemoveShippingAddressParams struct {
	Params
}

func (list *RemoveShippingAddressParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RemoveShippingAddress Remove an account's shipping address
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_shipping_address
//
// Returns: Shipping address deleted.
func (c *Client) RemoveShippingAddress(accountId string, shippingAddressId string, params *RemoveShippingAddressParams) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountSubscriptionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	// - When `state=active`, `state=canceled`, `state=expired`, or `state=future`, subscriptions with states that match the query and only those subscriptions will be returned.
	// - When `state=in_trial`, only subscriptions that have a trial_started_at date earlier than now and a trial_ends_at date later than now will be returned.
	// - When `state=live`, only subscriptions that are in an active, canceled, or future state or are in trial will be returned.
	State *string
}

func (list *ListAccountSubscriptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListAccountSubscriptionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListAccountSubscriptions List an account's subscriptions
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_subscriptions
//
// Returns: A list of the account's subscriptions.
func (c *Client) ListAccountSubscriptions(accountId string, params *ListAccountSubscriptionsParams) (*SubscriptionList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/subscriptions", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewSubscriptionList(c, path, params), nil
}

type ListAccountTransactionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Type - Filter by type field. The value `payment` will return both `purchase` and `capture` transactions.
	Type *string

	// Success - Filter by success field.
	Success *string
}

func (list *ListAccountTransactionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListAccountTransactionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	if list.Success != nil {
		options = append(options, KeyValue{Key: "success", Value: *list.Success})
	}

	return options
}

// ListAccountTransactions List an account's transactions
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_transactions
//
// Returns: A list of the account's transactions.
func (c *Client) ListAccountTransactions(accountId string, params *ListAccountTransactionsParams) (*TransactionList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/transactions", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewTransactionList(c, path, params), nil
}

type ListChildAccountsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Email - Filter for accounts with this exact email address. A blank value will return accounts with both `null` and `""` email addresses. Note that multiple accounts can share one email address.
	Email *string

	// Subscriber - Filter for accounts with or without a subscription in the `active`,
	// `canceled`, or `future` state.
	Subscriber *bool

	// PastDue - Filter for accounts with an invoice in the `past_due` state.
	PastDue *string
}

func (list *ListChildAccountsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListChildAccountsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Email != nil {
		options = append(options, KeyValue{Key: "email", Value: *list.Email})
	}

	if list.Subscriber != nil {
		options = append(options, KeyValue{Key: "subscriber", Value: strconv.FormatBool(*list.Subscriber)})
	}

	if list.PastDue != nil {
		options = append(options, KeyValue{Key: "past_due", Value: *list.PastDue})
	}

	return options
}

// ListChildAccounts List an account's child accounts
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_child_accounts
//
// Returns: A list of an account's child accounts.
func (c *Client) ListChildAccounts(accountId string, params *ListChildAccountsParams) (*AccountList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/accounts", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewAccountList(c, path, params), nil
}

type ListAccountAcquisitionParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListAccountAcquisitionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListAccountAcquisitionParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListAccountAcquisition List a site's account acquisition data
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_acquisition
//
// Returns: A list of the site's account acquisition data.
func (c *Client) ListAccountAcquisition(params *ListAccountAcquisitionParams) (*AccountAcquisitionList, error) {
	path, err := c.InterpolatePath("/acquisitions")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewAccountAcquisitionList(c, path, params), nil
}

type ListCouponsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListCouponsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListCouponsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListCoupons List a site's coupons
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_coupons
//
// Returns: A list of the site's coupons.
func (c *Client) ListCoupons(params *ListCouponsParams) (*CouponList, error) {
	path, err := c.InterpolatePath("/coupons")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewCouponList(c, path, params), nil
}

type CreateCouponParams struct {
	Params
}

func (list *CreateCouponParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateCoupon Create a new coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_coupon
//
// Returns: A new coupon.
func (c *Client) CreateCoupon(body *CouponCreate, params *CreateCouponParams) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons")
	if err != nil {
		return nil, err
	}
	result := &Coupon{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetCouponParams struct {
	Params
}

func (list *GetCouponParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetCoupon Fetch a coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_coupon
//
// Returns: A coupon.
func (c *Client) GetCoupon(couponId string, params *GetCouponParams) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	if err != nil {
		return nil, err
	}
	result := &Coupon{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdateCouponParams struct {
	Params
}

func (list *UpdateCouponParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdateCoupon Update an active coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_coupon
//
// Returns: The updated coupon.
func (c *Client) UpdateCoupon(couponId string, body *CouponUpdate, params *UpdateCouponParams) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	if err != nil {
		return nil, err
	}
	result := &Coupon{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type DeactivateCouponParams struct {
	Params
}

func (list *DeactivateCouponParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// DeactivateCoupon Expire a coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/deactivate_coupon
//
// Returns: The expired Coupon
func (c *Client) DeactivateCoupon(couponId string, params *DeactivateCouponParams) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	if err != nil {
		return nil, err
	}
	result := &Coupon{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RestoreCouponParams struct {
	Params
}

func (list *RestoreCouponParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RestoreCoupon Restore an inactive coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/restore_coupon
//
// Returns: The restored coupon.
func (c *Client) RestoreCoupon(couponId string, body *CouponUpdate, params *RestoreCouponParams) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}/restore", couponId)
	if err != nil {
		return nil, err
	}
	result := &Coupon{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListUniqueCouponCodesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListUniqueCouponCodesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListUniqueCouponCodesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListUniqueCouponCodes List unique coupon codes associated with a bulk coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_unique_coupon_codes
//
// Returns: A list of unique coupon codes that were generated
func (c *Client) ListUniqueCouponCodes(couponId string, params *ListUniqueCouponCodesParams) (*UniqueCouponCodeList, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}/unique_coupon_codes", couponId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewUniqueCouponCodeList(c, path, params), nil
}

type ListCreditPaymentsParams struct {
	Params

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListCreditPaymentsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListCreditPaymentsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListCreditPayments List a site's credit payments
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_credit_payments
//
// Returns: A list of the site's credit payments.
func (c *Client) ListCreditPayments(params *ListCreditPaymentsParams) (*CreditPaymentList, error) {
	path, err := c.InterpolatePath("/credit_payments")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewCreditPaymentList(c, path, params), nil
}

type GetCreditPaymentParams struct {
	Params
}

func (list *GetCreditPaymentParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetCreditPayment Fetch a credit payment
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_credit_payment
//
// Returns: A credit payment.
func (c *Client) GetCreditPayment(creditPaymentId string, params *GetCreditPaymentParams) (*CreditPayment, error) {
	path, err := c.InterpolatePath("/credit_payments/{credit_payment_id}", creditPaymentId)
	if err != nil {
		return nil, err
	}
	result := &CreditPayment{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListCustomFieldDefinitionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// RelatedType - Filter by related type.
	RelatedType *string
}

func (list *ListCustomFieldDefinitionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListCustomFieldDefinitionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.RelatedType != nil {
		options = append(options, KeyValue{Key: "related_type", Value: *list.RelatedType})
	}

	return options
}

// ListCustomFieldDefinitions List a site's custom field definitions
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_custom_field_definitions
//
// Returns: A list of the site's custom field definitions.
func (c *Client) ListCustomFieldDefinitions(params *ListCustomFieldDefinitionsParams) (*CustomFieldDefinitionList, error) {
	path, err := c.InterpolatePath("/custom_field_definitions")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewCustomFieldDefinitionList(c, path, params), nil
}

type GetCustomFieldDefinitionParams struct {
	Params
}

func (list *GetCustomFieldDefinitionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetCustomFieldDefinition Fetch an custom field definition
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_custom_field_definition
//
// Returns: An custom field definition.
func (c *Client) GetCustomFieldDefinition(customFieldDefinitionId string, params *GetCustomFieldDefinitionParams) (*CustomFieldDefinition, error) {
	path, err := c.InterpolatePath("/custom_field_definitions/{custom_field_definition_id}", customFieldDefinitionId)
	if err != nil {
		return nil, err
	}
	result := &CustomFieldDefinition{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListItemsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	State *string
}

func (list *ListItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListItemsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListItems List a site's items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_items
//
// Returns: A list of the site's items.
func (c *Client) ListItems(params *ListItemsParams) (*ItemList, error) {
	path, err := c.InterpolatePath("/items")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewItemList(c, path, params), nil
}

type CreateItemParams struct {
	Params
}

func (list *CreateItemParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateItem Create a new item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_item
//
// Returns: A new item.
func (c *Client) CreateItem(body *ItemCreate, params *CreateItemParams) (*Item, error) {
	path, err := c.InterpolatePath("/items")
	if err != nil {
		return nil, err
	}
	result := &Item{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetItemParams struct {
	Params
}

func (list *GetItemParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetItem Fetch an item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_item
//
// Returns: An item.
func (c *Client) GetItem(itemId string, params *GetItemParams) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}", itemId)
	if err != nil {
		return nil, err
	}
	result := &Item{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdateItemParams struct {
	Params
}

func (list *UpdateItemParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdateItem Update an active item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_item
//
// Returns: The updated item.
func (c *Client) UpdateItem(itemId string, body *ItemUpdate, params *UpdateItemParams) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}", itemId)
	if err != nil {
		return nil, err
	}
	result := &Item{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type DeactivateItemParams struct {
	Params
}

func (list *DeactivateItemParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// DeactivateItem Deactivate an item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/deactivate_item
//
// Returns: An item.
func (c *Client) DeactivateItem(itemId string, params *DeactivateItemParams) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}", itemId)
	if err != nil {
		return nil, err
	}
	result := &Item{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ReactivateItemParams struct {
	Params
}

func (list *ReactivateItemParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// ReactivateItem Reactivate an inactive item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/reactivate_item
//
// Returns: An item.
func (c *Client) ReactivateItem(itemId string, params *ReactivateItemParams) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}/reactivate", itemId)
	if err != nil {
		return nil, err
	}
	result := &Item{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListMeasuredUnitParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	State *string
}

func (list *ListMeasuredUnitParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListMeasuredUnitParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListMeasuredUnit List a site's measured units
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_measured_unit
//
// Returns: A list of the site's measured units.
func (c *Client) ListMeasuredUnit(params *ListMeasuredUnitParams) (*MeasuredUnitList, error) {
	path, err := c.InterpolatePath("/measured_units")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewMeasuredUnitList(c, path, params), nil
}

type CreateMeasuredUnitParams struct {
	Params
}

func (list *CreateMeasuredUnitParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateMeasuredUnit Create a new measured unit
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_measured_unit
//
// Returns: A new measured unit.
func (c *Client) CreateMeasuredUnit(body *MeasuredUnitCreate, params *CreateMeasuredUnitParams) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units")
	if err != nil {
		return nil, err
	}
	result := &MeasuredUnit{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetMeasuredUnitParams struct {
	Params
}

func (list *GetMeasuredUnitParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetMeasuredUnit Fetch a measured unit
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_measured_unit
//
// Returns: An item.
func (c *Client) GetMeasuredUnit(measuredUnitId string, params *GetMeasuredUnitParams) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units/{measured_unit_id}", measuredUnitId)
	if err != nil {
		return nil, err
	}
	result := &MeasuredUnit{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdateMeasuredUnitParams struct {
	Params
}

func (list *UpdateMeasuredUnitParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdateMeasuredUnit Update a measured unit
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_measured_unit
//
// Returns: The updated measured_unit.
func (c *Client) UpdateMeasuredUnit(measuredUnitId string, body *MeasuredUnitUpdate, params *UpdateMeasuredUnitParams) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units/{measured_unit_id}", measuredUnitId)
	if err != nil {
		return nil, err
	}
	result := &MeasuredUnit{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RemoveMeasuredUnitParams struct {
	Params
}

func (list *RemoveMeasuredUnitParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RemoveMeasuredUnit Remove a measured unit
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_measured_unit
//
// Returns: A measured unit.
func (c *Client) RemoveMeasuredUnit(measuredUnitId string, params *RemoveMeasuredUnitParams) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units/{measured_unit_id}", measuredUnitId)
	if err != nil {
		return nil, err
	}
	result := &MeasuredUnit{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListInvoicesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Type - Filter by type when:
	// - `type=charge`, only charge invoices will be returned.
	// - `type=credit`, only credit invoices will be returned.
	// - `type=non-legacy`, only charge and credit invoices will be returned.
	// - `type=legacy`, only legacy invoices will be returned.
	Type *string
}

func (list *ListInvoicesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListInvoicesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListInvoices List a site's invoices
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_invoices
//
// Returns: A list of the site's invoices.
func (c *Client) ListInvoices(params *ListInvoicesParams) (*InvoiceList, error) {
	path, err := c.InterpolatePath("/invoices")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewInvoiceList(c, path, params), nil
}

type GetInvoiceParams struct {
	Params
}

func (list *GetInvoiceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetInvoice Fetch an invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_invoice
//
// Returns: An invoice.
func (c *Client) GetInvoice(invoiceId string, params *GetInvoiceParams) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type PutInvoiceParams struct {
	Params
}

func (list *PutInvoiceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// PutInvoice Update an invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/put_invoice
//
// Returns: An invoice.
func (c *Client) PutInvoice(invoiceId string, body *InvoiceUpdatable, params *PutInvoiceParams) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type CollectInvoiceParams struct {
	Params

	// Body - The body of the request.
	Body *InvoiceCollect
}

func (list *CollectInvoiceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CollectInvoice Collect a pending or past due, automatic invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/collect_invoice
//
// Returns: The updated invoice.
func (c *Client) CollectInvoice(invoiceId string, params *CollectInvoiceParams) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/collect", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, params, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type FailInvoiceParams struct {
	Params
}

func (list *FailInvoiceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// FailInvoice Mark an open invoice as failed
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/fail_invoice
//
// Returns: The updated invoice.
func (c *Client) FailInvoice(invoiceId string, params *FailInvoiceParams) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/mark_failed", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type MarkInvoiceSuccessfulParams struct {
	Params
}

func (list *MarkInvoiceSuccessfulParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// MarkInvoiceSuccessful Mark an open invoice as successful
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/mark_invoice_successful
//
// Returns: The updated invoice.
func (c *Client) MarkInvoiceSuccessful(invoiceId string, params *MarkInvoiceSuccessfulParams) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/mark_successful", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ReopenInvoiceParams struct {
	Params
}

func (list *ReopenInvoiceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// ReopenInvoice Reopen a closed, manual invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/reopen_invoice
//
// Returns: The updated invoice.
func (c *Client) ReopenInvoice(invoiceId string, params *ReopenInvoiceParams) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/reopen", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type VoidInvoiceParams struct {
	Params
}

func (list *VoidInvoiceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// VoidInvoice Void a credit invoice.
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/void_invoice
//
// Returns: The updated invoice.
func (c *Client) VoidInvoice(invoiceId string, params *VoidInvoiceParams) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/void", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RecordExternalTransactionParams struct {
	Params
}

func (list *RecordExternalTransactionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RecordExternalTransaction Record an external payment for a manual invoices.
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/record_external_transaction
//
// Returns: The recorded transaction.
func (c *Client) RecordExternalTransaction(invoiceId string, body *ExternalTransaction, params *RecordExternalTransactionParams) (*Transaction, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/transactions", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Transaction{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListInvoiceLineItemsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Original - Filter by original field.
	Original *string

	// State - Filter by state field.
	State *string

	// Type - Filter by type field.
	Type *string
}

func (list *ListInvoiceLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListInvoiceLineItemsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Original != nil {
		options = append(options, KeyValue{Key: "original", Value: *list.Original})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListInvoiceLineItems List an invoice's line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_invoice_line_items
//
// Returns: A list of the invoice's line items.
func (c *Client) ListInvoiceLineItems(invoiceId string, params *ListInvoiceLineItemsParams) (*LineItemList, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/line_items", invoiceId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewLineItemList(c, path, params), nil
}

type ListInvoiceCouponRedemptionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListInvoiceCouponRedemptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListInvoiceCouponRedemptionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListInvoiceCouponRedemptions Show the coupon redemptions applied to an invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_invoice_coupon_redemptions
//
// Returns: A list of the the coupon redemptions associated with the invoice.
func (c *Client) ListInvoiceCouponRedemptions(invoiceId string, params *ListInvoiceCouponRedemptionsParams) (*CouponRedemptionList, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/coupon_redemptions", invoiceId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewCouponRedemptionList(c, path, params), nil
}

type ListRelatedInvoicesParams struct {
	Params
}

func (list *ListRelatedInvoicesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// ListRelatedInvoices List an invoice's related credit or charge invoices
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_related_invoices
//
// Returns: A list of the credit or charge invoices associated with the invoice.
func (c *Client) ListRelatedInvoices(invoiceId string, params *ListRelatedInvoicesParams) (*InvoiceList, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/related_invoices", invoiceId)
	if err != nil {
		return nil, err
	}
	params.IsListOp = true
	return NewInvoiceList(c, path, params), nil
}

type RefundInvoiceParams struct {
	Params
}

func (list *RefundInvoiceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RefundInvoice Refund an invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/refund_invoice
//
// Returns: Returns the new credit invoice.
func (c *Client) RefundInvoice(invoiceId string, body *InvoiceRefund, params *RefundInvoiceParams) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/refund", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListLineItemsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Original - Filter by original field.
	Original *string

	// State - Filter by state field.
	State *string

	// Type - Filter by type field.
	Type *string
}

func (list *ListLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListLineItemsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Original != nil {
		options = append(options, KeyValue{Key: "original", Value: *list.Original})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListLineItems List a site's line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_line_items
//
// Returns: A list of the site's line items.
func (c *Client) ListLineItems(params *ListLineItemsParams) (*LineItemList, error) {
	path, err := c.InterpolatePath("/line_items")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewLineItemList(c, path, params), nil
}

type GetLineItemParams struct {
	Params
}

func (list *GetLineItemParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetLineItem Fetch a line item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_line_item
//
// Returns: A line item.
func (c *Client) GetLineItem(lineItemId string, params *GetLineItemParams) (*LineItem, error) {
	path, err := c.InterpolatePath("/line_items/{line_item_id}", lineItemId)
	if err != nil {
		return nil, err
	}
	result := &LineItem{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RemoveLineItemParams struct {
	Params
}

func (list *RemoveLineItemParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RemoveLineItem Delete an uninvoiced line item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_line_item
//
// Returns: Line item deleted.
func (c *Client) RemoveLineItem(lineItemId string, params *RemoveLineItemParams) (*Empty, error) {
	path, err := c.InterpolatePath("/line_items/{line_item_id}", lineItemId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListPlansParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	State *string
}

func (list *ListPlansParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListPlansParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListPlans List a site's plans
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_plans
//
// Returns: A list of plans.
func (c *Client) ListPlans(params *ListPlansParams) (*PlanList, error) {
	path, err := c.InterpolatePath("/plans")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewPlanList(c, path, params), nil
}

type CreatePlanParams struct {
	Params
}

func (list *CreatePlanParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreatePlan Create a plan
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_plan
//
// Returns: A plan.
func (c *Client) CreatePlan(body *PlanCreate, params *CreatePlanParams) (*Plan, error) {
	path, err := c.InterpolatePath("/plans")
	if err != nil {
		return nil, err
	}
	result := &Plan{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetPlanParams struct {
	Params
}

func (list *GetPlanParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetPlan Fetch a plan
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_plan
//
// Returns: A plan.
func (c *Client) GetPlan(planId string, params *GetPlanParams) (*Plan, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}", planId)
	if err != nil {
		return nil, err
	}
	result := &Plan{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdatePlanParams struct {
	Params
}

func (list *UpdatePlanParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdatePlan Update a plan
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_plan
//
// Returns: A plan.
func (c *Client) UpdatePlan(planId string, body *PlanUpdate, params *UpdatePlanParams) (*Plan, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}", planId)
	if err != nil {
		return nil, err
	}
	result := &Plan{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RemovePlanParams struct {
	Params
}

func (list *RemovePlanParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RemovePlan Remove a plan
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_plan
//
// Returns: Plan deleted
func (c *Client) RemovePlan(planId string, params *RemovePlanParams) (*Plan, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}", planId)
	if err != nil {
		return nil, err
	}
	result := &Plan{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListPlanAddOnsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	State *string
}

func (list *ListPlanAddOnsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListPlanAddOnsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListPlanAddOns List a plan's add-ons
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_plan_add_ons
//
// Returns: A list of add-ons.
func (c *Client) ListPlanAddOns(planId string, params *ListPlanAddOnsParams) (*AddOnList, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons", planId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewAddOnList(c, path, params), nil
}

type CreatePlanAddOnParams struct {
	Params
}

func (list *CreatePlanAddOnParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreatePlanAddOn Create an add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_plan_add_on
//
// Returns: An add-on.
func (c *Client) CreatePlanAddOn(planId string, body *AddOnCreate, params *CreatePlanAddOnParams) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons", planId)
	if err != nil {
		return nil, err
	}
	result := &AddOn{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetPlanAddOnParams struct {
	Params
}

func (list *GetPlanAddOnParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetPlanAddOn Fetch a plan's add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_plan_add_on
//
// Returns: An add-on.
func (c *Client) GetPlanAddOn(planId string, addOnId string, params *GetPlanAddOnParams) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	if err != nil {
		return nil, err
	}
	result := &AddOn{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdatePlanAddOnParams struct {
	Params
}

func (list *UpdatePlanAddOnParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdatePlanAddOn Update an add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_plan_add_on
//
// Returns: An add-on.
func (c *Client) UpdatePlanAddOn(planId string, addOnId string, body *AddOnUpdate, params *UpdatePlanAddOnParams) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	if err != nil {
		return nil, err
	}
	result := &AddOn{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RemovePlanAddOnParams struct {
	Params
}

func (list *RemovePlanAddOnParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RemovePlanAddOn Remove an add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_plan_add_on
//
// Returns: Add-on deleted
func (c *Client) RemovePlanAddOn(planId string, addOnId string, params *RemovePlanAddOnParams) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	if err != nil {
		return nil, err
	}
	result := &AddOn{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAddOnsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	State *string
}

func (list *ListAddOnsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListAddOnsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListAddOns List a site's add-ons
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_add_ons
//
// Returns: A list of add-ons.
func (c *Client) ListAddOns(params *ListAddOnsParams) (*AddOnList, error) {
	path, err := c.InterpolatePath("/add_ons")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewAddOnList(c, path, params), nil
}

type GetAddOnParams struct {
	Params
}

func (list *GetAddOnParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetAddOn Fetch an add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_add_on
//
// Returns: An add-on.
func (c *Client) GetAddOn(addOnId string, params *GetAddOnParams) (*AddOn, error) {
	path, err := c.InterpolatePath("/add_ons/{add_on_id}", addOnId)
	if err != nil {
		return nil, err
	}
	result := &AddOn{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListShippingMethodsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListShippingMethodsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListShippingMethodsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListShippingMethods List a site's shipping methods
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_shipping_methods
//
// Returns: A list of the site's shipping methods.
func (c *Client) ListShippingMethods(params *ListShippingMethodsParams) (*ShippingMethodList, error) {
	path, err := c.InterpolatePath("/shipping_methods")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewShippingMethodList(c, path, params), nil
}

type CreateShippingMethodParams struct {
	Params
}

func (list *CreateShippingMethodParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateShippingMethod Create a new shipping method
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_shipping_method
//
// Returns: A new shipping method.
func (c *Client) CreateShippingMethod(body *ShippingMethodCreate, params *CreateShippingMethodParams) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods")
	if err != nil {
		return nil, err
	}
	result := &ShippingMethod{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetShippingMethodParams struct {
	Params
}

func (list *GetShippingMethodParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetShippingMethod Fetch a shipping method
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_shipping_method
//
// Returns: A shipping method.
func (c *Client) GetShippingMethod(shippingMethodId string, params *GetShippingMethodParams) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods/{shipping_method_id}", shippingMethodId)
	if err != nil {
		return nil, err
	}
	result := &ShippingMethod{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdateShippingMethodParams struct {
	Params
}

func (list *UpdateShippingMethodParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdateShippingMethod Update an active Shipping Method
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_shipping_method
//
// Returns: The updated shipping method.
func (c *Client) UpdateShippingMethod(shippingMethodId string, body *ShippingMethodUpdate, params *UpdateShippingMethodParams) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods/{shipping_method_id}", shippingMethodId)
	if err != nil {
		return nil, err
	}
	result := &ShippingMethod{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type DeactivateShippingMethodParams struct {
	Params
}

func (list *DeactivateShippingMethodParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// DeactivateShippingMethod Deactivate a shipping method
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/deactivate_shipping_method
//
// Returns: A shipping method.
func (c *Client) DeactivateShippingMethod(shippingMethodId string, params *DeactivateShippingMethodParams) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods/{shipping_method_id}", shippingMethodId)
	if err != nil {
		return nil, err
	}
	result := &ShippingMethod{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListSubscriptionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	// - When `state=active`, `state=canceled`, `state=expired`, or `state=future`, subscriptions with states that match the query and only those subscriptions will be returned.
	// - When `state=in_trial`, only subscriptions that have a trial_started_at date earlier than now and a trial_ends_at date later than now will be returned.
	// - When `state=live`, only subscriptions that are in an active, canceled, or future state or are in trial will be returned.
	State *string
}

func (list *ListSubscriptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListSubscriptionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListSubscriptions List a site's subscriptions
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_subscriptions
//
// Returns: A list of the site's subscriptions.
func (c *Client) ListSubscriptions(params *ListSubscriptionsParams) (*SubscriptionList, error) {
	path, err := c.InterpolatePath("/subscriptions")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewSubscriptionList(c, path, params), nil
}

type CreateSubscriptionParams struct {
	Params
}

func (list *CreateSubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateSubscription Create a new subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_subscription
//
// Returns: A subscription.
func (c *Client) CreateSubscription(body *SubscriptionCreate, params *CreateSubscriptionParams) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions")
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetSubscriptionParams struct {
	Params
}

func (list *GetSubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetSubscription Fetch a subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_subscription
//
// Returns: A subscription.
func (c *Client) GetSubscription(subscriptionId string, params *GetSubscriptionParams) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ModifySubscriptionParams struct {
	Params
}

func (list *ModifySubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// ModifySubscription Modify a subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/modify_subscription
//
// Returns: A subscription.
func (c *Client) ModifySubscription(subscriptionId string, body *SubscriptionUpdate, params *ModifySubscriptionParams) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type TerminateSubscriptionParams struct {
	Params

	// Refund - The type of refund to perform:
	// * `full` - Performs a full refund of the last invoice for the current subscription term.
	// * `partial` - Prorates a refund based on the amount of time remaining in the current bill cycle.
	// * `none` - Terminates the subscription without a refund.
	// In the event that the most recent invoice is a $0 invoice paid entirely by credit, Recurly will apply the credit back to the customer’s account.
	// You may also terminate a subscription with no refund and then manually refund specific invoices.
	Refund *string
}

func (list *TerminateSubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *TerminateSubscriptionParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Refund != nil {
		options = append(options, KeyValue{Key: "refund", Value: *list.Refund})
	}

	return options
}

// TerminateSubscription Terminate a subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/terminate_subscription
//
// Returns: An expired subscription.
func (c *Client) TerminateSubscription(subscriptionId string, params *TerminateSubscriptionParams) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodDelete, path, params, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type CancelSubscriptionParams struct {
	Params

	// Body - The body of the request.
	Body *SubscriptionCancel
}

func (list *CancelSubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CancelSubscription Cancel a subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/cancel_subscription
//
// Returns: A canceled or failed subscription.
func (c *Client) CancelSubscription(subscriptionId string, params *CancelSubscriptionParams) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/cancel", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, params, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ReactivateSubscriptionParams struct {
	Params
}

func (list *ReactivateSubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// ReactivateSubscription Reactivate a canceled subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/reactivate_subscription
//
// Returns: An active subscription.
func (c *Client) ReactivateSubscription(subscriptionId string, params *ReactivateSubscriptionParams) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/reactivate", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type PauseSubscriptionParams struct {
	Params
}

func (list *PauseSubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// PauseSubscription Pause subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/pause_subscription
//
// Returns: A subscription.
func (c *Client) PauseSubscription(subscriptionId string, body *SubscriptionPause, params *PauseSubscriptionParams) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/pause", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ResumeSubscriptionParams struct {
	Params
}

func (list *ResumeSubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// ResumeSubscription Resume subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/resume_subscription
//
// Returns: A subscription.
func (c *Client) ResumeSubscription(subscriptionId string, params *ResumeSubscriptionParams) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/resume", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ConvertTrialParams struct {
	Params
}

func (list *ConvertTrialParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// ConvertTrial Convert trial subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/convert_trial
//
// Returns: A subscription.
func (c *Client) ConvertTrial(subscriptionId string, params *ConvertTrialParams) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/convert_trial", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetSubscriptionChangeParams struct {
	Params
}

func (list *GetSubscriptionChangeParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetSubscriptionChange Fetch a subscription's pending change
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_subscription_change
//
// Returns: A subscription's pending change.
func (c *Client) GetSubscriptionChange(subscriptionId string, params *GetSubscriptionChangeParams) (*SubscriptionChange, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &SubscriptionChange{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type CreateSubscriptionChangeParams struct {
	Params
}

func (list *CreateSubscriptionChangeParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateSubscriptionChange Create a new subscription change
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_subscription_change
//
// Returns: A subscription change.
func (c *Client) CreateSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate, params *CreateSubscriptionChangeParams) (*SubscriptionChange, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &SubscriptionChange{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RemoveSubscriptionChangeParams struct {
	Params
}

func (list *RemoveSubscriptionChangeParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RemoveSubscriptionChange Delete the pending subscription change
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_subscription_change
//
// Returns: Subscription change was deleted.
func (c *Client) RemoveSubscriptionChange(subscriptionId string, params *RemoveSubscriptionChangeParams) (*Empty, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type PreviewSubscriptionChangeParams struct {
	Params
}

func (list *PreviewSubscriptionChangeParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// PreviewSubscriptionChange Preview a new subscription change
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/preview_subscription_change
//
// Returns: A subscription change.
func (c *Client) PreviewSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate, params *PreviewSubscriptionChangeParams) (*SubscriptionChange, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change/preview", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &SubscriptionChange{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListSubscriptionInvoicesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Type - Filter by type when:
	// - `type=charge`, only charge invoices will be returned.
	// - `type=credit`, only credit invoices will be returned.
	// - `type=non-legacy`, only charge and credit invoices will be returned.
	// - `type=legacy`, only legacy invoices will be returned.
	Type *string
}

func (list *ListSubscriptionInvoicesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListSubscriptionInvoicesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListSubscriptionInvoices List a subscription's invoices
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_subscription_invoices
//
// Returns: A list of the subscription's invoices.
func (c *Client) ListSubscriptionInvoices(subscriptionId string, params *ListSubscriptionInvoicesParams) (*InvoiceList, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/invoices", subscriptionId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewInvoiceList(c, path, params), nil
}

type ListSubscriptionLineItemsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Original - Filter by original field.
	Original *string

	// State - Filter by state field.
	State *string

	// Type - Filter by type field.
	Type *string
}

func (list *ListSubscriptionLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListSubscriptionLineItemsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Original != nil {
		options = append(options, KeyValue{Key: "original", Value: *list.Original})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListSubscriptionLineItems List a subscription's line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_subscription_line_items
//
// Returns: A list of the subscription's line items.
func (c *Client) ListSubscriptionLineItems(subscriptionId string, params *ListSubscriptionLineItemsParams) (*LineItemList, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/line_items", subscriptionId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewLineItemList(c, path, params), nil
}

type ListSubscriptionCouponRedemptionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListSubscriptionCouponRedemptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListSubscriptionCouponRedemptionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListSubscriptionCouponRedemptions Show the coupon redemptions for a subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_subscription_coupon_redemptions
//
// Returns: A list of the the coupon redemptions on a subscription.
func (c *Client) ListSubscriptionCouponRedemptions(subscriptionId string, params *ListSubscriptionCouponRedemptionsParams) (*CouponRedemptionList, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/coupon_redemptions", subscriptionId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewCouponRedemptionList(c, path, params), nil
}

type ListUsageParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `usage_timestamp` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=usage_timestamp` or `sort=recorded_timestamp`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=usage_timestamp` or `sort=recorded_timestamp`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// BillingStatus - Filter by usage record's billing status
	BillingStatus *string
}

func (list *ListUsageParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListUsageParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.BillingStatus != nil {
		options = append(options, KeyValue{Key: "billing_status", Value: *list.BillingStatus})
	}

	return options
}

// ListUsage List a subscription add-on's usage records
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_usage
//
// Returns: A list of the subscription add-on's usage records.
func (c *Client) ListUsage(subscriptionId string, addOnId string, params *ListUsageParams) (*UsageList, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/add_ons/{add_on_id}/usage", subscriptionId, addOnId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewUsageList(c, path, params), nil
}

type CreateUsageParams struct {
	Params
}

func (list *CreateUsageParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreateUsage Log a usage record on this subscription add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_usage
//
// Returns: The created usage record.
func (c *Client) CreateUsage(subscriptionId string, addOnId string, body *UsageCreate, params *CreateUsageParams) (*Usage, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/add_ons/{add_on_id}/usage", subscriptionId, addOnId)
	if err != nil {
		return nil, err
	}
	result := &Usage{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetUsageParams struct {
	Params
}

func (list *GetUsageParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetUsage Get a usage record
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_usage
//
// Returns: The usage record.
func (c *Client) GetUsage(usageId string, params *GetUsageParams) (*Usage, error) {
	path, err := c.InterpolatePath("/usage/{usage_id}", usageId)
	if err != nil {
		return nil, err
	}
	result := &Usage{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type UpdateUsageParams struct {
	Params
}

func (list *UpdateUsageParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// UpdateUsage Update a usage record
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_usage
//
// Returns: The updated usage record.
func (c *Client) UpdateUsage(usageId string, body *UsageCreate, params *UpdateUsageParams) (*Usage, error) {
	path, err := c.InterpolatePath("/usage/{usage_id}", usageId)
	if err != nil {
		return nil, err
	}
	result := &Usage{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type RemoveUsageParams struct {
	Params
}

func (list *RemoveUsageParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// RemoveUsage Delete a usage record.
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_usage
//
// Returns: Usage was successfully deleted.
func (c *Client) RemoveUsage(usageId string, params *RemoveUsageParams) (*Empty, error) {
	path, err := c.InterpolatePath("/usage/{usage_id}", usageId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListTransactionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Type - Filter by type field. The value `payment` will return both `purchase` and `capture` transactions.
	Type *string

	// Success - Filter by success field.
	Success *string
}

func (list *ListTransactionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

func (list *ListTransactionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	if list.Success != nil {
		options = append(options, KeyValue{Key: "success", Value: *list.Success})
	}

	return options
}

// ListTransactions List a site's transactions
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_transactions
//
// Returns: A list of the site's transactions.
func (c *Client) ListTransactions(params *ListTransactionsParams) (*TransactionList, error) {
	path, err := c.InterpolatePath("/transactions")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	params.IsListOp = true
	return NewTransactionList(c, path, params), nil
}

type GetTransactionParams struct {
	Params
}

func (list *GetTransactionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetTransaction Fetch a transaction
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_transaction
//
// Returns: A transaction.
func (c *Client) GetTransaction(transactionId string, params *GetTransactionParams) (*Transaction, error) {
	path, err := c.InterpolatePath("/transactions/{transaction_id}", transactionId)
	if err != nil {
		return nil, err
	}
	result := &Transaction{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetUniqueCouponCodeParams struct {
	Params
}

func (list *GetUniqueCouponCodeParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetUniqueCouponCode Fetch a unique coupon code
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_unique_coupon_code
//
// Returns: A unique coupon code.
func (c *Client) GetUniqueCouponCode(uniqueCouponCodeId string, params *GetUniqueCouponCodeParams) (*UniqueCouponCode, error) {
	path, err := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}", uniqueCouponCodeId)
	if err != nil {
		return nil, err
	}
	result := &UniqueCouponCode{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type DeactivateUniqueCouponCodeParams struct {
	Params
}

func (list *DeactivateUniqueCouponCodeParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// DeactivateUniqueCouponCode Deactivate a unique coupon code
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/deactivate_unique_coupon_code
//
// Returns: A unique coupon code.
func (c *Client) DeactivateUniqueCouponCode(uniqueCouponCodeId string, params *DeactivateUniqueCouponCodeParams) (*UniqueCouponCode, error) {
	path, err := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}", uniqueCouponCodeId)
	if err != nil {
		return nil, err
	}
	result := &UniqueCouponCode{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ReactivateUniqueCouponCodeParams struct {
	Params
}

func (list *ReactivateUniqueCouponCodeParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// ReactivateUniqueCouponCode Restore a unique coupon code
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/reactivate_unique_coupon_code
//
// Returns: A unique coupon code.
func (c *Client) ReactivateUniqueCouponCode(uniqueCouponCodeId string, params *ReactivateUniqueCouponCodeParams) (*UniqueCouponCode, error) {
	path, err := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}/restore", uniqueCouponCodeId)
	if err != nil {
		return nil, err
	}
	result := &UniqueCouponCode{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type CreatePurchaseParams struct {
	Params
}

func (list *CreatePurchaseParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// CreatePurchase Create a new purchase
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_purchase
//
// Returns: Returns the new invoices
func (c *Client) CreatePurchase(body *PurchaseCreate, params *CreatePurchaseParams) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/purchases")
	if err != nil {
		return nil, err
	}
	result := &InvoiceCollection{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type PreviewPurchaseParams struct {
	Params
}

func (list *PreviewPurchaseParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// PreviewPurchase Preview a new purchase
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/preview_purchase
//
// Returns: Returns preview of the new invoices
func (c *Client) PreviewPurchase(body *PurchaseCreate, params *PreviewPurchaseParams) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/purchases/preview")
	if err != nil {
		return nil, err
	}
	result := &InvoiceCollection{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetExportDatesParams struct {
	Params
}

func (list *GetExportDatesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetExportDates List the dates that have an available export to download.
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_export_dates
//
// Returns: Returns a list of dates.
func (c *Client) GetExportDates(params *GetExportDatesParams) (*ExportDates, error) {
	path, err := c.InterpolatePath("/export_dates")
	if err != nil {
		return nil, err
	}
	result := &ExportDates{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type GetExportFilesParams struct {
	Params
}

func (list *GetExportFilesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		IsListOp:       list.IsListOp,
		RequestParams:  list,
	}
}

// GetExportFiles List of the export files that are available to download.
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_export_files
//
// Returns: Returns a list of export files to download.
func (c *Client) GetExportFiles(exportDate string, params *GetExportFilesParams) (*ExportFiles, error) {
	path, err := c.InterpolatePath("/export_dates/{export_date}/export_files", exportDate)
	if err != nil {
		return nil, err
	}
	result := &ExportFiles{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}
