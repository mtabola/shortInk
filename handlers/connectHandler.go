package handlers

import (
	"../functions"
	"../globalVars"
	"net/http"
	"strings"
)

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
	hash := strings.Replace(r.URL.Path, "/show/", "", 1)

	connectLink := globalVars.Links.FindLinkByShortHash(hash)

	if connectLink == nil {
		w = functions.GetResponse("Undefined shortLink", "Invalid link"+hash, http.StatusBadRequest, w)
		return
	}

	http.Redirect(w, r, connectLink.FullLink, http.StatusSeeOther)
}
