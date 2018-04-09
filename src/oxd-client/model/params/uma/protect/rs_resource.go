//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//

package protect

// Resource structure
type RsResource struct {
	Path string `json:"path"`
	Conditions []Condition `json:"conditions"`
}
