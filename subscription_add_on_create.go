// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type SubscriptionAddOnCreate struct {

	// If `add_on_source` is set to `plan_add_on` or left blank, then plan's add-on `code` should be used.
	// If `add_on_source` is set to `item`, then the `code` from the associated item should be used.
	Code *string `json:"code,omitempty"`

	// Used to determine where the associated add-on data is pulled from. If this value is set to
	// `plan_add_on` or left blank, then add-on data will be pulled from the plan's add-ons. If the associated
	// `plan` has `allow_any_item_on_subscriptions` set to `true` and this field is set to `item`, then
	// the associated add-on data will be pulled from the site's item catalog.
	AddOnSource *string `json:"add_on_source,omitempty"`

	// Quantity
	Quantity *int `json:"quantity,omitempty"`

	// Allows up to 2 decimal places. Optionally, override the add-on's default unit amount.
	// If the plan add-on's `tier_type` is `tiered`, `volume`, or `stairstep`, then `unit_amount` cannot be provided.
	UnitAmount *float64 `json:"unit_amount,omitempty"`

	// Allows up to 9 decimal places.  Optionally, override the add-on's default unit amount.
	// If the plan add-on's `tier_type` is `tiered`, `volume`, or `stairstep`, then `unit_amount_decimal` cannot be provided.
	// Only supported when the plan add-on's `add_on_type` = `usage`.
	// If `unit_amount_decimal` is provided, `unit_amount` cannot be provided.
	UnitAmountDecimal *string `json:"unit_amount_decimal,omitempty"`

	// If the plan add-on's `tier_type` is `flat`, then `tiers` must be absent. The `tiers` object
	// must include one to many tiers with `ending_quantity` and `unit_amount`.
	// There must be one tier with an `ending_quantity` of 999999999 which is the
	// default if not provided. See our [Guide](https://developers.recurly.com/guides/item-addon-guide.html)
	// for an overview of how to configure quantity-based pricing models.
	Tiers []SubscriptionAddOnTierCreate `json:"tiers,omitempty"`

	// The percentage taken of the monetary amount of usage tracked. This can be up to 4 decimal places. A value between 0.0 and 100.0. Required if `add_on_type` is usage and `usage_type` is percentage. Must be omitted otherwise. `usage_percentage` does not support tiers. See our [Guide](https://developers.recurly.com/guides/usage-based-billing-guide.html) for an overview of how to configure usage add-ons.
	UsagePercentage *float64 `json:"usage_percentage,omitempty"`

	// Revenue schedule type
	RevenueScheduleType *string `json:"revenue_schedule_type,omitempty"`
}
