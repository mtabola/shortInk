package handles

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"../globalVars"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	url := strings.Replace(r.URL.Path, "/", "", 1)

	if url == "" {
		t, err := template.ParseFiles(ComDir + "header.html", ComDir + "index.html", ComDir + "footer.html")

		if err != nil {
			fmt.Fprintf(w, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.ExecuteTemplate(w, "index", globalVars.Links.Links)
	}
}

