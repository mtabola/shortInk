package handles

import (
	"net/http"
	"strconv"

	"../globalVars"
)

func DeleteHandle(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
	}

	numId, _ := strconv.Atoi(id)

	for _, link := range globalVars.Links.Links {
		if link.LinkId == numId {
			err := globalVars.Links.DeleteLink(link, globalVars.DB)

			if err != nil {
				globalVars.Response.GetResponse(err.Error(), err.Error())
				http.Redirect(w, r, "/response", http.StatusBadRequest)
			}
		}
	}

	globalVars.Response.GetResponse("Link is deleted", nil)

	http.Redirect(w, r, "/response", http.StatusFound)
}
