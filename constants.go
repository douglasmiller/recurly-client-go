package recurly

import (
	"encoding/json"
	"errors"
	"fmt"
)

type RelatedType int

const (
	UnsetEnum = iota
	UndefinedEnum
	AccountEnum
	ItemEnum
	SubscriptionEnum
	BananaEnum
)

func (e RelatedType) String() string {
	names := [...]string{"", "Undefined", "account", "item", "subscription"}
	if e == 0 || e > SubscriptionEnum {
		return ""
	}
	return names[e]
}

// MarshalJSON marshals the enum as a quoted json string
func (e RelatedType) MarshalJSON() ([]byte, error) {
	if e.String() == "" {
		return nil, errors.New("Unknown RelatedType enum")
	}
	return []byte(fmt.Sprintf("\"%s\"", e)), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (e *RelatedType) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*e = map[string]RelatedType{"account": AccountEnum, "item": ItemEnum, "subscription": SubscriptionEnum}[j]
	// If we don't have an enum for the value, use Undefined
	if *e == UnsetEnum {
		*e = UndefinedEnum
	}
	return nil
}
