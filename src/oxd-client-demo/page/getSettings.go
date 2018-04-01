package page

import (
	"oxd-client-demo/conf"
	"net/http"
	"encoding/json"
	"io"
	"oxd-client-demo/utils"
)

func GetSettingsSite(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session conf.SessionVars )  {
	w.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(configuration)
	io.WriteString(w, string(result))
}

func PostSettingsSite(w http.ResponseWriter, r *http.Request, configuration *conf.Configuration, session conf.SessionVars )  {
	utils.ParseResponse(w,r,configuration)
}
