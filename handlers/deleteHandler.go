package handlers

import (
	"../functions"
	"../globalVars"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	hash, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w = functions.GetResponse("Can't read body", err.Error(), http.StatusBadRequest, w)
		return
	}

	//srtHash := strings.ReplaceAll(string(hash), "\r\n", "")
	//srtHash = strings.ReplaceAll(srtHash, " ", "")

	delLink := globalVars.Links.FindLinkByShortHash(string(hash))
	if delLink == nil {
		w = functions.GetResponse("This link not exist", "", http.StatusBadRequest, w)
		return
	}

	err = globalVars.Links.DeleteLink(*delLink, globalVars.DB)

	if err != nil {
		w = functions.GetResponse("Link is not deleted", err.Error(), http.StatusBadRequest, w)
		return
	}

	w = functions.GetResponse("Link is deleted", "", http.StatusOK, w)
}
