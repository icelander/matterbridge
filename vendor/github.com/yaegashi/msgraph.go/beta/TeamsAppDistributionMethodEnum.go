// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TeamsAppDistributionMethod undocumented
type TeamsAppDistributionMethod int

const (
	// TeamsAppDistributionMethodVStore undocumented
	TeamsAppDistributionMethodVStore TeamsAppDistributionMethod = 0
	// TeamsAppDistributionMethodVOrganization undocumented
	TeamsAppDistributionMethodVOrganization TeamsAppDistributionMethod = 1
	// TeamsAppDistributionMethodVSideloaded undocumented
	TeamsAppDistributionMethodVSideloaded TeamsAppDistributionMethod = 2
	// TeamsAppDistributionMethodVUnknownFutureValue undocumented
	TeamsAppDistributionMethodVUnknownFutureValue TeamsAppDistributionMethod = 3
)

// TeamsAppDistributionMethodPStore returns a pointer to TeamsAppDistributionMethodVStore
func TeamsAppDistributionMethodPStore() *TeamsAppDistributionMethod {
	v := TeamsAppDistributionMethodVStore
	return &v
}

// TeamsAppDistributionMethodPOrganization returns a pointer to TeamsAppDistributionMethodVOrganization
func TeamsAppDistributionMethodPOrganization() *TeamsAppDistributionMethod {
	v := TeamsAppDistributionMethodVOrganization
	return &v
}

// TeamsAppDistributionMethodPSideloaded returns a pointer to TeamsAppDistributionMethodVSideloaded
func TeamsAppDistributionMethodPSideloaded() *TeamsAppDistributionMethod {
	v := TeamsAppDistributionMethodVSideloaded
	return &v
}

// TeamsAppDistributionMethodPUnknownFutureValue returns a pointer to TeamsAppDistributionMethodVUnknownFutureValue
func TeamsAppDistributionMethodPUnknownFutureValue() *TeamsAppDistributionMethod {
	v := TeamsAppDistributionMethodVUnknownFutureValue
	return &v
}