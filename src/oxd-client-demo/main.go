//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package main

import (
	"net/http"
	//"oxd-client-demo/page"
	//"oxd-client-demo/conf"
	"log"
	"github.com/juju/loggo"
	"github.com/BurntSushi/toml"
	"oxd-client/utils"
	"strings"
	"oxd-client-demo/conf"
	"oxd-client-demo/page"
	"encoding/json"
)



var serverConf conf.Configuration
var session conf.SessionVars
var globalvariables conf.GlobalVars
var oxdSettingsInterface interface{}
//func LoadoxdSetting(){
//
//	ReadoxdSettings,_ := ioutil.ReadFile("src/oxd-client-demo/utils/oxd_config.json")
//
//	json.Unmarshal(ReadoxdSettings,&oxdSettingsInterface)
//	oxdSettingValues := oxdSettingsInterface.(map[string]interface{})
//
//	globalvariables.Httpresturl = oxdSettingValues["https_rest_url"].(string)
//	globalvariables.ConnectionType = oxdSettingValues["connection_type"].(string)
//	globalvariables.ClientName = oxdSettingValues["client_name"].(string)
//    globalvariables.RedirectUri = oxdSettingValues["authorization_redirect_uri"].(string)
//	globalvariables.LogoutUri = oxdSettingValues["post_logout_redirect_uri"].(string)
//	globalvariables.ClientName = oxdSettingValues["client_name"].(string)
//	globalvariables.Clientid = oxdSettingValues["client_id"].(string)
//	globalvariables.Clientsecret = oxdSettingValues["client_secret"].(string)
//	globalvariables.Ophost = oxdSettingValues["op_host"].(string)
//	globalvariables.Oxdid = oxdSettingValues["oxd_id"].(string)
//	globalvariables.OxdPort = oxdSettingValues["oxd_host_port"].(string)
//	globalvariables.ConnectionType = oxdSettingValues["connection_type"].(string)
//	globalvariables.OxdHost = oxdSettingValues["oxd_host"].(string)
//
//	ScopeInterfaceType := oxdSettingValues["scope"].([]interface{})
//	ScopeStringArray := make([]string, len(ScopeInterfaceType))
//	for i := range ScopeInterfaceType {
//		ScopeStringArray[i] = ScopeInterfaceType[i].(string)
//	}
//	globalvariables.Scope = ScopeStringArray
//	GrantTypeInterfaceType := oxdSettingValues["grant_types"].([]interface{})
//	GrantTypeStringArray := make([]string, len(GrantTypeInterfaceType))
//	for i := range GrantTypeInterfaceType {
//		GrantTypeStringArray[i] = GrantTypeInterfaceType[i].(string)
//	}
//	globalvariables.GrantTypes = GrantTypeStringArray
//	globalvariables.Host = fmt.Sprint(globalvariables.OxdHost+":"+globalvariables.OxdPort)
//}
//func GetTemplate(name string) *template.Template{
//	tmpl := template.Must(template.ParseFiles("src/oxd-client-demo/templates/" + name + ".html",
//	))
//	return tmpl
//}
//var isdynamicresponse struct {
//	Response  string            `json:"response"`
//}
func main() {

	//http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
	//	LoadoxdSetting()
	//	page.SetupClientPage(w, r, serverConf, &session,globalvariables)
	//})
	//
	//http.HandleFunc("/checkregistration", func(w http.ResponseWriter, r *http.Request) {
	//
	//   isdynamic := page.IsDynamicRegistrationPage(w, r, serverConf, &session)
	//   isdynamicresponse.Response = isdynamic
	//   IsdynamicResponseJson, _ := json.Marshal(isdynamicresponse)
	//   w.Write(IsdynamicResponseJson)
	//
	//})
	//
	//http.HandleFunc("/loadoxdsettings", func(w http.ResponseWriter, r *http.Request) {
	//	LoadoxdSetting()
	//	 w.Header().Set("Content-Type", "application/json")
	//	 loadsettingsjson, _ := json.Marshal(globalvariables)
	//     w.Write(loadsettingsjson)
	//
	//})
	//
	//
	//http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
	//	LoadoxdSetting()
	//	ProtectionAccessToken := page.GetClientTokenPageSite(serverConf)
	//	page.UpdateSitePage(w, r, serverConf, &session , ProtectionAccessToken, globalvariables)
	//})
	//
	//http.HandleFunc("/authUrl", func(w http.ResponseWriter, r *http.Request) {
	//	LoadoxdSetting()
	//	ProtectionAccessToken := page.GetClientTokenPageSite(serverConf)
	//	page.AuthorizationUrlPageSite(w, r, serverConf,session,ProtectionAccessToken,globalvariables)
	//
	//})
	//
	//http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
	//	LoadoxdSetting()
	//	ProtectionAccessToken := page.GetClientTokenPageSite(serverConf)
	//	GetUserInfo := page.RedirectPage(w, r, serverConf, &session , ProtectionAccessToken,globalvariables )
	//	err := GetTemplate("userinfo").Execute(w,GetUserInfo)
	//	if err != nil {
     //       panic(err)
     //   }
	//
	//})
	//
	//http.HandleFunc("/userInfo", func(w http.ResponseWriter, r *http.Request) {
	//	//page.UserInfoPage(w, r, serverConf, session)
	//
	//
	//})
	//
	//http.HandleFunc("/authCodeFlow", func(w http.ResponseWriter, r *http.Request) {
	//	page.AuthorizationCodeFlowPageSite(w, r, serverConf,session)
	//})
	//
	//http.HandleFunc("/authCode", func(w http.ResponseWriter, r *http.Request) {
	//	page.AuthorizationCodePageSite(w, r, serverConf, &session)
	//})
	//
	//http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
	//	page.ValidationPageSite(w, r, serverConf, session)
	//})
	//
	//http.HandleFunc("/logoutUrl", func(w http.ResponseWriter, r *http.Request) {
	//	LoadoxdSetting()
	//	ProtectionAccessToken := page.GetClientTokenPageSite(serverConf)
	//	page.LogoutUrlPageSite(w, r, serverConf, session, ProtectionAccessToken,globalvariables)
	//})
	//
	//http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
	//	page.LogoutPageSite(w, r)
	//})
	//
	//http.HandleFunc("/protectresources", func(w http.ResponseWriter, r *http.Request) {
	//	LoadoxdSetting()
	//
	//	ProtectionAccessToken := page.GetClientTokenPageSite(serverConf)
	//	page.ProtectResourcesPage(w, r, serverConf, &session, ProtectionAccessToken,globalvariables)
	//})
	//
	//http.HandleFunc("/checkaccess", func(w http.ResponseWriter, r *http.Request) {
	//	LoadoxdSetting()
	//	ProtectionAccessToken := page.GetClientTokenPageSite(serverConf)
	//	page.CheckAccessPage(w, r, serverConf, &session, ProtectionAccessToken,globalvariables)
	//})
	//
	//http.HandleFunc("/getrpt", func(w http.ResponseWriter, r *http.Request) {
	//	LoadoxdSetting()
	//	ProtectionAccessToken := page.GetClientTokenPageSite(serverConf)
	//	page.RpGetRptPage(w, r, serverConf, &session, ProtectionAccessToken,globalvariables)
	//})
	//http.HandleFunc("/getclaims", func(w http.ResponseWriter, r *http.Request) {
	//	LoadoxdSetting()
	//	ProtectionAccessToken := page.GetClientTokenPageSite(serverConf)
	//	page.RpGetClaimsGatheringUrlPage(w, r, serverConf, &session, ProtectionAccessToken,globalvariables)
	//})
	//http.HandleFunc("/deletesettings", func(w http.ResponseWriter, r *http.Request) {
	//	jsonio.DeleteoxdSetting()
	//})


	http.HandleFunc("/getAccessToken", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(page.GetClientTokenPageSite(&serverConf)))
	})

	http.HandleFunc("/registerSite", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(page.RegisterClientSite(&serverConf)))
	})

	http.HandleFunc("/updateSite", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(page.UpdateClientSite(w,r,&serverConf)))
	})

	http.HandleFunc("/setupClient", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(page.SetupClientPage(w,r,&serverConf)))
	})

	http.HandleFunc("/removeSite", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(page.RemoveClientSite(&serverConf)))
	})

	http.HandleFunc("/authorizationUrl", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(page.GetAuthorizationUrlPage(&serverConf)))
	})

	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		page.ReadCode(w,r,&serverConf)
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(page.GetTokenWithCode(&serverConf)))
	})

	http.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(page.GetTokenWithRefreshToken(&serverConf)))
	})

	http.HandleFunc("/userInfo", func(w http.ResponseWriter, r *http.Request) {
		data,_ :=json.Marshal(page.GetUserInfo(&serverConf))
		w.Write(data)
	})

	http.HandleFunc("/logoutUrl", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(page.GetLogoutUrl(&serverConf)))
	})

	http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			page.GetSettingsSite(w, r, serverConf, session)
		case "POST":
			page.PostSettingsSite(w, r, &serverConf, session)
		default:
			return
		}
	})

	http.Handle("/", http.FileServer(http.Dir("src/oxd-client-demo/resources/")))
	   err := http.ListenAndServeTLS(":8080",serverConf.Cert, serverConf.Key, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}

	func init() {
		//LoadoxdSetting()
		 _, err := toml.DecodeFile("src/oxd-client-demo/conf/main.toml", &serverConf)
		 utils.CheckError("transport.transport","Config file error newone",err)

		if strings.EqualFold(serverConf.Loglevel, "Debug") {
			loggo.GetLogger("default").SetLogLevel(loggo.DEBUG)
		}
	}

