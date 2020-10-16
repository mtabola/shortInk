package handles

import (
	"fmt"
	"net/http"
	"html/template"

	"../globalVars"
)

func ResponseHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(ComDir + "header.html", ComDir + "response.html", ComDir + "footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t.ExecuteTemplate(w, "response", globalVars.Response)
}
