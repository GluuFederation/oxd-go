//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package validation
// Deprecated:
type LicenseStatusRequestParams struct {
	Features string `json:"features,omitempty"`
}
// Deprecated:
type LicenseStatusResponseParams struct {
	Valid bool `json:"valid"`
	ThreadCount int64 `json:"thread_count"`
	Name string `json:"name"`
	Features []string `json:"features"`
}
