//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package main

import (
	"net/http"
	"log"
	"github.com/juju/loggo"
	"github.com/BurntSushi/toml"
	"oxd-client/utils"
	"strings"
	"oxd-client-demo/conf"
	"oxd-client-demo/service"
	"encoding/json"
)



var serverConf conf.Configuration
func main() {
	http.HandleFunc("/getAccessToken", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.GetClientTokenPageSite(&serverConf)))
	})

	http.HandleFunc("/registerSite", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.RegisterClientSite(&serverConf)))
	})

	http.HandleFunc("/updateSite", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.UpdateClientSite(w,r,&serverConf)))
	})

	http.HandleFunc("/setupClient", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.SetupClientPage(w,r,&serverConf)))
	})

	http.HandleFunc("/removeSite", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.RemoveClientSite(&serverConf)))
	})

	http.HandleFunc("/authorizationUrl", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.GetAuthorizationUrlPage(&serverConf)))
	})

	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		service.ReadCode(w,r,&serverConf)
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.GetTokenWithCode(&serverConf)))
	})

	http.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.GetTokenWithRefreshToken(&serverConf)))
	})

	http.HandleFunc("/userInfo", func(w http.ResponseWriter, r *http.Request) {
		data,_ :=json.Marshal(service.GetUserInfo(&serverConf))
		w.Write(data)
	})

	http.HandleFunc("/logoutUrl", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.GetLogoutUrl(&serverConf)))
	})

	http.HandleFunc("/umaRsProtect", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.UmaRsProtect(&serverConf)))
	})

	http.HandleFunc("/umaCheckAccess", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.UmaCheckAccess(&serverConf)))
	})

	http.HandleFunc("/umaRpt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.UmaGetRpt(&serverConf)))
	})

	http.HandleFunc("/umaClaimsUrl", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(service.UmaGetClaimsGathering(&serverConf)))
	})

	http.HandleFunc("/umaIntrospectRpt", func(w http.ResponseWriter, r *http.Request) {
		data,_ :=json.Marshal(service.GetUserInfo(&serverConf))
		w.Write(data)
	})

	http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			service.GetSettingsSite(w, r, serverConf)
		case "POST":
			service.PostSettingsSite(w, r, &serverConf)
		default:
			return
		}
	})

	http.Handle("/", http.FileServer(http.Dir("src/oxd-client-demo/webapp/")))
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

