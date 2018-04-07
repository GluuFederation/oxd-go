package uma

// Get Claims gathering request https://gluu.org/docs/oxd/3.1.2/api/#uma-rp-get-claims-gathering-url
type RpGetClaimsGatheringUrlRequestParams struct {
	OxdId string `json:"oxd_id"`
	Ticket string `json:"ticket"`
	ClaimsRedirectURI string `json:"claims_redirect_uri"`
	ProtectionAccessToken string `json:"protection_access_token"`
}

// Get Claims gathering response https://gluu.org/docs/oxd/3.1.2/api/#uma-rp-get-claims-gathering-url
type RpGetClaimsGatheringUrlResponseParams struct {
	URL string `json:"url"`
	State string `json:"state"`
	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`
}