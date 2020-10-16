package handles

import (
	"fmt"
	"net/http"
	"html/template"
	"strconv"

	"../globalVars"
	"../structures"
)

func EditHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(ComDir + "header.html", ComDir + "linkInterection.html", ComDir + "footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")

	numId, err := strconv.Atoi(id)

	var link structures.Link
	if err == nil {
		link = globalVars.Links.Links[numId]
	} else {
		http.NotFound(w, r)
	}


	t.ExecuteTemplate(w, "linkInterection", link)
}
