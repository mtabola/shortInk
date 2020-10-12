package handlers

import (
	"../functions"
	"fmt"
	"net/http"
	"strings"
	"../globalVars"
	"../structures"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)

	inFLink := r.FormValue("fullLink")
	inSLink := r.FormValue("shortLink")
	/*strBody := string(body)
	var custLink []string*/

	inFLink = strings.ReplaceAll(inFLink, " ", "")
	if (!strings.HasPrefix(inFLink, "http://") && !strings.HasPrefix(inFLink, "https://")) || len(strings.Split(inFLink, ".")) != 2 {
		w = functions.GetResponse("Link is not correct", "", http.StatusBadRequest, w)
		return
	}


	fmt.Printf("%s, %s", inFLink, inSLink)



	if(inSLink == "") {
		inSLink = fmt.Sprintf("%x", functions.HashGeneration(inFLink))
	}

	lksLen := len(globalVars.Links.Links)

	var id int
	if lksLen != 0 {
		id = (globalVars.Links.Links[(lksLen - 1)].LinkId) + 1
	} else {
		id = 1
	}

	newLink := structures.Link{
		LinkId:    id,
		FullLink:  inFLink,
		ShortLink: inSLink,
	}

	err := globalVars.Links.AddLink(newLink, globalVars.DB)

	if err != nil {
		w = functions.GetResponse("AddLink fault", err.Error(), http.StatusBadRequest, w)
		return
	}

	w = functions.GetResponse("New link added", "", http.StatusOK, w)
}
