//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package protect

type Condition struct {
	HttpMethods []string `json:"httpMethods"`
	Scopes []string `json:"scopes"`
	TicketScopes []string `json:"ticketScopes"`
	ScopeExpression *ScopeExpression `json:"scope_expression,omitempty"`
}

type ScopeExpression struct {
	Rule interface{}`json:"rule"`
	Data []string `json:"data"`
}