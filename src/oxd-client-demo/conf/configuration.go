//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package conf

import "oxd-client/model/params/uma/protect"


type Configuration struct{

	IdpHost string `toml:"idpHost"`
	Host string `toml:"host"`
	OxdHttpsHost string `toml:"oxdHttpsHost"`
	Key string `toml:"key"`
	Cert string `toml:"cert"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	ClientId string `toml:"clientId"`
	ClientSecret string `toml:"clientSecret"`
	ClientRegistrationClientUri string `toml:"clientRegistrationClientUri"`
	Loglevel string `toml:"loglevel"`
	Path string `toml:"path"`
	Type string `toml:"type"`
	Condition []protect.Condition `toml:"condition"`
	AccessToken string
	SetupClient bool
	ClientOxdId string

	RegisterSiteParams RegisterSiteParams `toml:"registerSiteParams"`
	OpenIdConnect OpenIdConnect
	Uma Uma `toml:"uma"`
}


type Uma struct {
	Resources []UmaResources `toml:"resources"`
	ResourcesOxdId string
	CheckAccessPath string `toml:"checkAccessPath"`
	CheckAccessMethod string `toml:"checkAccessMethod"`
	Ticket string
	Rpt string
	ClaimsUrl string
	IntrospectRpt interface{}
}

type UmaResources struct {
	Path string `toml:"path"`
	Conditions []UmaConditions `toml:"conditions"`
}

type UmaConditions struct {
	HttpMethods []string `toml:"httpMethods"`
	Scopes []string `toml:"scopes"`
	TicketScopes []string `toml:"ticketScopes"`
}

type OpenIdConnect struct {
	AuthorizationUrl string
	Code string
	State string
	IdToken string
	AccessToken string
	RefreshToken string
	Claims map[string] []string
	LogoutUrl string
}

type RegisterSiteParams struct{
	OxdId string
	OpHost string  `toml:"opHost"`
	AuthorizationRedirectUri string  `toml:"authorizationRedirectUri"`
	PostLogoutRedirectUri string  `toml:"postLogoutRedirectUri"`
	ApplicationType string  `toml:"applicationType"`

	RedirectUris []string  `toml:"redirectUris"`
	ResponseTypes []string  `toml:"responseTypes"`

	ClientId string  `toml:"clientId"`
	ClientSecret string  `toml:"clientSecret"`
	ClientName string  `toml:"clientName"`
	ClientJwksUri string  `toml:"clientJwksUri"`
	ClientTokenEndpointAuthMethod string  `toml:"clientTokenEndpointAuthMethod"`
	ClientRequestUris  []string  `toml:"clientRequestUris"`
	ClientLogoutUri  []string  `toml:"clientLogoutUris"`
	ClientSectorIdentifierUri string  `toml:"clientSectorIdentifierUri"`
	ClientFrontChannelUris []string  `toml:"client_frontchannelLogoutUris"`
	ClientRegistrationAccessToken string  `toml:"clientRegistrationAccessToken"`
	ClientRegistrationClientUri string  `toml:"clientRegistrationClientUri"`

	Scope  []string  `toml:"scope"`
	UiLocales  []string  `toml:"uiLocales"`
	ClaimsLocales  []string  `toml:"claimsLocales"`
	ClaimsRedirectUri  []string  `toml:"claimsRedirectUri"`
	AcrValues  []string  `toml:"acrValues"`
	GrantTypes  []string  `toml:"grantType"`
	Contacts []string  `toml:"contacts"`
}
