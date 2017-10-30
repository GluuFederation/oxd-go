package page



import (
	"fmt"
	"io/ioutil"
	"net/http"
    "oxd-client-demo/conf"
	"crypto/tls"
	"encoding/json"
)

func IsDynamicRegistrationPage(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session *conf.SessionVars) (string)   {
	var Response = ""
	ophost := r.FormValue("OphostId")
	openidconnecturl := fmt.Sprint(ophost+"/.well-known/openid-configuration")
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	
	client := &http.Client{Transport: tr}
	resp, err := client.Get(openidconnecturl)
	
	if err != nil {
		panic(err)
	}
	
	defer resp.Body.Close()
	
	ReadConfigurationResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var ConfigurationString interface{}
	json.Unmarshal(ReadConfigurationResponse, &ConfigurationString)
	configurationValues := ConfigurationString.(map[string]interface{})
	if(configurationValues["registration_endpoint"] == nil) {
		
		Response = "false"

	}
	if(configurationValues["registration_endpoint"] != nil) {
		
		Response = "true"
		
	}
	
    return Response
	
}