package tests

import (
	"testing"
	"oxd-client-test/utils"
	"oxd-client/model/params/uma"
	"oxd-client/constants"
	"oxd-client/client"
	"github.com/stretchr/testify/assert"
	"oxd-client-test/conf"
	"oxd-client/model/transport"
)

func TestSocketRpGetClaimsGatheringUrl(t *testing.T) {
	executeRpGetClaimsGatheringUrlTest(t,utils.GetSocketRequest)
}

func TestRestRpGetClaimsGatheringUrl(t *testing.T) {
	executeRpGetClaimsGatheringUrlTest(t,utils.GetRestRequest)
}

func executeRpGetClaimsGatheringUrlTest(t *testing.T,getRequest utils.GetRequest) {
	//BEFORE
	setupParams := utils.SetupClient()
	utils.ProtectRs(setupParams.OxdId, getRequest)
	checkAccessParams := utils.CheckAccess(setupParams.OxdId, "", "/ws/phone", "GET", getRequest)
	requestParams := uma.RpGetClaimsGatheringUrlRequestParams{setupParams.OxdId,checkAccessParams.Ticket,
	conf.TestConfiguration.RedirectUrl,""}

	request,connectionParam := getRequest(constants.RP_GET_CLAIMS_GATHERING_URL,requestParams)
	var response transport.OxdResponse
	var responseParams uma.RpGetClaimsGatheringUrlResponseParams

	//TEST
	client.Send(request,connectionParam,&response)

	//ASSERT
	response.GetParams(&responseParams)
	assert.Equal(t,constants.STATUS_OK,response.Status,"Status should be ok")
	assert.NotEmpty(t,responseParams.URL,"Url should not be empty")
	assert.NotEmpty(t,responseParams.State,"State should not be empty")
}
