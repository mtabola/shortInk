package handlers

import (
	"../globalVars"
	"encoding/json"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	LinksJson, _ := json.Marshal(globalVars.Links.GetAllLinks())
	w.WriteHeader(http.StatusOK)
	w.Write(LinksJson)
}
