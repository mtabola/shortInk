package handles

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"../globalVars"
	"../structures"
	"../functions"
)

func SaveHandle(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	fullLink :=r.FormValue("fullLink")
	shortLink :=r.FormValue("shortLink")

	var l *structures.Link
	var editStatus bool = true

	if id == "" {
		lksLen := len(globalVars.Links.Links)
		if lksLen != 0 {
			id = strconv.Itoa((globalVars.Links.Links[(lksLen - 1)].LinkId) + 1)
		} else {
			id = "1"
		}
		editStatus = false
	}
	numId, err := strconv.Atoi(id)

	if err != nil {
		globalVars.Response.GetResponse("Operation fault", err.Error())
		goto Redirect
	}

	if shortLink == "" {
		shortLink = fmt.Sprintf("%x", functions.HashGeneration(fullLink))
	}

	shortLink = strings.ReplaceAll(shortLink, " ", "-")

	l = structures.NewLink(numId, fullLink, shortLink)

	for i, link := range globalVars.Links.Links {
		if link.LinkId == l.LinkId && editStatus{
			globalVars.Links.Links[i] = *l
			err := globalVars.Links.EditLink(*l, globalVars.DB)

			if err != nil {
				globalVars.Response.GetResponse(err.Error(), err.Error())
				goto Redirect
			}
			globalVars.Response.GetResponse("Link is edited", nil)
			goto Redirect
		}
	}

	err = globalVars.Links.AddLink(*l, globalVars.DB)

	if err != nil {
		globalVars.Response.GetResponse("Operation fault", err.Error())
	} else {
		globalVars.Response.GetResponse("Link is added", nil)
	}

	Redirect:
	http.Redirect(w, r, "/response", http.StatusFound)
}