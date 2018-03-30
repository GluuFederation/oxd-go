package page

import (
	"oxd-client-demo/conf"
	"net/http"
	"encoding/json"
	"io"
	"fmt"
)

func GetSettingsSite(w http.ResponseWriter, r *http.Request, configuration conf.Configuration, session conf.SessionVars )  {
	w.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(configuration)
	io.WriteString(w, string(result))
}

func PostSettingsSite(w http.ResponseWriter, r *http.Request, configuration *conf.Configuration, session conf.SessionVars )  {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(configuration)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
}
