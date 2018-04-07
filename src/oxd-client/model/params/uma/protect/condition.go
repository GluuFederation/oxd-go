//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//

//Package protect contains structures for protect resources call
package protect

// Condition structure
type Condition struct {
	HttpMethods []string `json:"httpMethods"`
	Scopes []string `json:"scopes"`
	TicketScopes []string `json:"ticketScopes"`
	ScopeExpression *ScopeExpression `json:"scope_expression,omitempty"`
}

// Scope expression structure. Read more https://gluu.org/docs/ce/3.1.2/admin-guide/uma.md
type ScopeExpression struct {
	Rule interface{}`json:"rule"`
	Data []string `json:"data"`
}