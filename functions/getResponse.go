package functions

import (
	"../structures"
	"encoding/json"
	"net/http"
)

func GetResponse(msg string, err string, reqStatus int, w http.ResponseWriter) http.ResponseWriter {
	resp := &structures.Response{
		Message: msg,
		Error:   err,
	}

	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
	w.WriteHeader(reqStatus)
	return w
}
