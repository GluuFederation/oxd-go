//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package utils

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func DisplayResponse(w http.ResponseWriter,v interface{}){
	value, _ := json.Marshal(v)
	fmt.Fprintln(w,string(value))
	fmt.Println(string(value))

}

func SaveoxdSetting(w http.ResponseWriter,m interface{}){
	
	SettingsValue, _ := json.Marshal(m)
	ioutil.WriteFile("src/oxd-client-demo/utils/oxd_config.json", SettingsValue, 0644)

}

func UpdateoxdSetting(w http.ResponseWriter,c interface{}){
	fmt.Println(c)
	raw,_ := ioutil.ReadFile("src/oxd-client-demo/utils/oxd_config.json")
	var data interface{}
	json.Unmarshal(raw,&data)
	
	newdata := data.(map[string]interface{})
	updatedata := c.(map[string]interface{})
	newdata["op_host"] = updatedata["op_host"]
	newdata["client_name"] = updatedata["client_name"]
	newdata["authorization_redirect_uri"] = updatedata["authorization_redirect_uri"]
    newdata["post_logout_redirect_uri"] = updatedata["post_logout_redirect_uri"]
	newdata["client_id"] = updatedata["client_id"]
	newdata["client_secret"] = updatedata["client_secret"]
	 
	 newdata["connection_type"] = updatedata["connection_type"]
	 newdata["oxd_host"] = updatedata["oxd_host"]

	NewValue, _ := json.Marshal(data)
	ioutil.WriteFile("src/oxd-client-demo/utils/oxd_config.json", NewValue, 0644)
	

}

