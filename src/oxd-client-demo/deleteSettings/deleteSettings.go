//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package jsonio

import (
	"encoding/json"
	"io/ioutil"
)

func DeleteoxdSetting(){
	raw,_ := ioutil.ReadFile("src/oxd-client-demo/utils/oxd_config.json")
	var data interface{}
	json.Unmarshal(raw,&data)
	
	newdata := data.(map[string]interface{})
	
	 newdata["oxd_id"] = " "
	 newdata["op_host"] = " "
	 newdata["oxd_host"] = " "
	 newdata["client_id"] = " "
	 newdata["client_name"] = " "
	 newdata["oxd_host_port"] = " "
	 newdata["client_secret"] = " "
	 newdata["https_rest_url"] = " "
	 newdata["connection_type"] = " "
	 newdata["post_logout_redirect_uri"] = " "
	 newdata["dynamic_registration"] = "false "
	 newdata["authorization_redirect_uri"] = " "
	 SettingsValue, _ := json.Marshal(data)
	 ioutil.WriteFile("src/oxd-client-demo/utils/oxd_config.json", SettingsValue, 0644)
	  

}