package handlers

import (
	"../functions"
	"../globalVars"
	"../structures"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w = functions.GetResponse("Can't read the request body", err.Error(), http.StatusBadRequest, w)
		return
	}
	strBody := string(body)
	var custLink []string

	strBody = strings.ReplaceAll(strBody, " ", "")

	if (!strings.HasPrefix(strBody, "http://") && !strings.HasPrefix(strBody, "https://")) || len(strings.Split(strBody, ".")) != 2 {
		w = functions.GetResponse("Link is not correct", "", http.StatusBadRequest, w)
		return
	}

	if !strings.Contains(strBody, "==>") {
		strBody = fmt.Sprintf("%s==>%x", strBody, functions.HashGeneration(strBody))
	}

	custLink = strings.Split(strBody, "==>")

	lksLen := len(globalVars.Links.Links)

	var id int
	if lksLen != 0 {
		id = (globalVars.Links.Links[(lksLen - 1)].LinkId) + 1
	} else {
		id = 1
	}

	newLink := structures.Link{
		LinkId:    id,
		FullLink:  custLink[0],
		ShortLink: custLink[1],
	}

	err = globalVars.Links.AddLink(newLink, globalVars.DB)

	if err != nil {
		w = functions.GetResponse("AddLink fault", err.Error(), http.StatusBadRequest, w)
		return
	}

	w = functions.GetResponse("New link added", "", http.StatusOK, w)
}
