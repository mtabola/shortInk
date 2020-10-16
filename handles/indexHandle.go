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
	} else {
		for _, link := range globalVars.Links.Links {
			if link.ShortLink == url {
				http.Redirect(w, r, link.FullLink, http.StatusFound)
				return
			}
		}
		globalVars.Response.GetResponse("Link not found", "Link not found")
		http.Redirect(w, r, "/response", http.StatusNotFound)
	}
}

