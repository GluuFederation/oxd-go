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

	RegisterSiteParams RegisterSiteParams `toml:"registerSiteParams"`
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
	GrantTypes  []string  `toml:"grantTypes"`
	Contacts []string  `toml:"contacts"`
}
