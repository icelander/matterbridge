// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// IdentityUserRisk undocumented
type IdentityUserRisk struct {
	// Object is the base model of IdentityUserRisk
	Object
	// Level undocumented
	Level *UserRiskLevel `json:"level,omitempty"`
	// LastChangedDateTime undocumented
	LastChangedDateTime *time.Time `json:"lastChangedDateTime,omitempty"`
}