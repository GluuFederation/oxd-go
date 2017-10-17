package uma


type RpGetClaimsGatheringUrlRequestParams struct {
	OxdId string `json:"oxd_id"`
	Ticket string `json:"ticket"`
	ClaimsRedirectURI string `json:"claims_redirect_uri"`
	ProtectionAccessToken string `json:"protection_access_token"`
}

type RpGetClaimsGatheringUrlResponseParams struct {
	URL string `json:"url"`
	State string `json:"state"`
	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`
}