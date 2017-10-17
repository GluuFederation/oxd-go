//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package main

import (
	
	"net/http"
	//"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"oxd-client-demo/page"
	"oxd-client-demo/conf"
	"log"
	"github.com/juju/loggo"
	"github.com/BurntSushi/toml"
	"oxd-client/utils"
	"oxd-client-demo/deleteSettings"
	"strings"
	"text/template"
	
)



var serverConf conf.Configuration
var session conf.SessionVars
var globalvariables conf.GlobalVars
var oxdSettingsInterface interface{}
func LoadoxdSetting(){
	
	ReadoxdSettings,_ := ioutil.ReadFile("src/oxd-client-demo/utils/oxd_config.json")
	
	json.Unmarshal(ReadoxdSettings,&oxdSettingsInterface)
	oxdSettingValues := oxdSettingsInterface.(map[string]interface{})

	globalvariables.Httpresturl = oxdSettingValues["https_rest_url"].(string)
	globalvariables.ClientName = oxdSettingValues["client_name"].(string)
    globalvariables.RedirectUri = oxdSettingValues["authorization_redirect_uri"].(string)
	globalvariables.LogoutUri = oxdSettingValues["post_logout_redirect_uri"].(string)
	globalvariables.ClientName = oxdSettingValues["client_name"].(string)
	globalvariables.Clientid = oxdSettingValues["client_id"].(string)
	globalvariables.Clientsecret = oxdSettingValues["client_secret"].(string)
	globalvariables.Ophost = oxdSettingValues["op_host"].(string)
	globalvariables.Oxdid = oxdSettingValues["oxd_id"].(string)
	globalvariables.OxdPort = oxdSettingValues["oxd_host_port"].(string)
	globalvariables.ConnectionType = oxdSettingValues["connection_type"].(string)
	globalvariables.OxdHost = oxdSettingValues["oxd_host"].(string)
	globalvariables.Host = fmt.Sprint(globalvariables.OxdHost+":"+globalvariables.OxdPort)
}
func GetTemplate(name string) *template.Template{
	tmpl := template.Must(template.ParseFiles("src/oxd-client-demo/templates/" + name + ".html",
	))
	return tmpl
}
func main() {

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		page.SetupClientPage(w, r, serverConf, &session)
		
	})

	http.HandleFunc("/loadoxdsettings", func(w http.ResponseWriter, r *http.Request) {
		LoadoxdSetting()
		 w.Header().Set("Content-Type", "application/json")
		 loadsettingsjson, _ := json.Marshal(globalvariables)
	     w.Write(loadsettingsjson)
		
	})


	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		LoadoxdSetting()
		ProtectionAccessToken := page.GetClientTokenPageSite(w, r, serverConf,session ,globalvariables)
		page.UpdateSitePage(w, r, serverConf, &session , ProtectionAccessToken, globalvariables)
	})

	http.HandleFunc("/authUrl", func(w http.ResponseWriter, r *http.Request) {
		LoadoxdSetting()
		ProtectionAccessToken := page.GetClientTokenPageSite(w, r, serverConf,session,globalvariables)
		page.AuthorizationUrlPageSite(w, r, serverConf,session,ProtectionAccessToken,globalvariables)
		
	})

	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		LoadoxdSetting()
		ProtectionAccessToken := page.GetClientTokenPageSite(w, r, serverConf,session,globalvariables)
		GetUserInfo := page.RedirectPage(w, r, serverConf, &session , ProtectionAccessToken,globalvariables )
		err := GetTemplate("userinfo").Execute(w,GetUserInfo)
		if err != nil {
            panic(err)
        }

	})

	http.HandleFunc("/userInfo", func(w http.ResponseWriter, r *http.Request) {
		//page.UserInfoPage(w, r, serverConf, session)
		

	})

	http.HandleFunc("/authCodeFlow", func(w http.ResponseWriter, r *http.Request) {
		page.AuthorizationCodeFlowPageSite(w, r, serverConf,session)
	})

	http.HandleFunc("/authCode", func(w http.ResponseWriter, r *http.Request) {
		page.AuthorizationCodePageSite(w, r, serverConf, &session)
	})

	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		page.ValidationPageSite(w, r, serverConf, session)
	})

	http.HandleFunc("/logoutUrl", func(w http.ResponseWriter, r *http.Request) {
		LoadoxdSetting()
		ProtectionAccessToken := page.GetClientTokenPageSite(w, r, serverConf,session,globalvariables)
		page.LogoutUrlPageSite(w, r, serverConf, session, ProtectionAccessToken,globalvariables)
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		page.LogoutPageSite(w, r)
	})

	http.HandleFunc("/protectresources", func(w http.ResponseWriter, r *http.Request) {
		LoadoxdSetting()
		ProtectionAccessToken := page.GetClientTokenPageSite(w, r, serverConf,session,globalvariables)
		page.ProtectResourcesPage(w, r, serverConf, &session, ProtectionAccessToken,globalvariables)
	})

	http.HandleFunc("/checkaccess", func(w http.ResponseWriter, r *http.Request) {
		LoadoxdSetting()
		ProtectionAccessToken := page.GetClientTokenPageSite(w, r, serverConf,session,globalvariables)
		page.CheckAccessPage(w, r, serverConf, &session, ProtectionAccessToken,globalvariables)
	})

	http.HandleFunc("/getrpt", func(w http.ResponseWriter, r *http.Request) {
		LoadoxdSetting()
		ProtectionAccessToken := page.GetClientTokenPageSite(w, r, serverConf,session,globalvariables)
		page.RpGetRptPage(w, r, serverConf, &session, ProtectionAccessToken,globalvariables)
	})
	http.HandleFunc("/getclaims", func(w http.ResponseWriter, r *http.Request) {
		LoadoxdSetting()
		ProtectionAccessToken := page.GetClientTokenPageSite(w, r, serverConf,session,globalvariables)
		page.RpGetClaimsGatheringUrlPage(w, r, serverConf, &session, ProtectionAccessToken,globalvariables)
	})
	http.HandleFunc("/deletesettings", func(w http.ResponseWriter, r *http.Request) {
		jsonio.DeleteoxdSetting()
	})

	http.Handle("/", http.FileServer(http.Dir("src/oxd-client-demo/resources/")))
	   err := http.ListenAndServeTLS(":8080",serverConf.Cert, serverConf.Key, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	
	}
	
	
	func init() {
		LoadoxdSetting()
		 _, err := toml.DecodeFile("src/oxd-client-demo/conf/main.toml", &serverConf)
		 utils.CheckError("transport.transport","Config file error newone",err)
	
		if strings.EqualFold(serverConf.Loglevel, "Debug") {
			loggo.GetLogger("default").SetLogLevel(loggo.DEBUG)
		}
	}
	
