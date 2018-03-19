package model

type SetupClientRequestParams struct {

	OpHost string  `json:"op_host"`
	AuthorizationRedirectUri string  `json:"authorization_redirect_uri"`
	PostLogoutRedirectUri string  `json:"post_logout_redirect_uri"`
	ApplicationType string `json:"application_type"`

	RedirectUris []string  `json:"redirect_uris"`
	ResponseTypes []string  `json:"response_types"`

	ClientId string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	ClientName string  `json:"client_name"`
	ClientJwksUri string  `json:"client_jwks_uri"`
	ClientTokenEndpointAuthMethod string  `json:"client_token_endpoint_auth_method,empty"`
	ClientRequestUris  []string  `json:"client_request_uris"`
	ClientLogoutUri  []string  `json:"client_logout_uris"`
	ClientSectorIdentifierUri string  `json:"client_sector_identifier_uri"`

	Scope  []string  `json:"scope,omitempty"`
	UiLocales  []string  `json:"ui_locales"`
	ClaimsLocales  []string  `json:"claims_locales"`
	AcrValues  []string  `json:"acr_values"`
	GrantTypes  []string  `json:"grant_types"`
	Contacts []string  `json:"contacts"`
}

type SetupClientResponseParams struct {
	OxdId string `json:"oxd_id"`
	OpHost string `json:"op_host"`
	ClientId string  `json:"client_id"`
	ClientSecret string  `json:"client_secret"`
	}