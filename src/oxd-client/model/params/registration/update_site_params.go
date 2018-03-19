//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package model

import "time"

type UpdateSiteRequestParams struct {
	OxdId                         string `json:"oxd_id"`
	AuthorizationRedirectUri      string `json:"authorization_redirect_uri"`
	PostLogoutRedirectUri         string  `json:"post_logout_redirect_uri"`
	ProtectionAccessToken         string `json:"protection_access_token,omitempty"`

	ResponseTypes                 []string  `json:"response_types,omitempty"`

	ClientId                    string  `json:"client_id"`
	ClientSecret                  string  `json:"client_secret"`
	ClientJwksUri                 string  `json:"client_jwks_uri"`
	ClientSectorIdentifierUri     string  `json:"client_sector_identifier_uri,omitempty"`
	ClientTokenEndpointAuthMethod string  `json:"client_token_endpoint_auth_method,omitempty"`
	ClientRequestUris             []string  `json:"client_request_uris,omitempty"`
	ClientSecretExpiresAt         time.Time  `json:"client_secret_expires_at,omitempty"`

	Scope                         []string  `json:"scope"`
	UiLocales                     []string  `json:"ui_locales,omitempty"`
	ClaimsLocales                 []string  `json:"claims_locales,omitempty"`
	AcrValues                     []string  `json:"acr_values"`
	GrantType                     []string  `json:"grant_types"`
	Contacts                      []string  `json:"contacts,omitempty"`
}

type UpdateSiteResponseParams struct {
	OxdId string `json:"oxd_id"`
}