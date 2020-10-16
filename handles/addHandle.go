package handles

import (
	"fmt"
	"html/template"
	"net/http"
)

func AddHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(ComDir + "header.html", ComDir + "linkInterection.html", ComDir + "footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t.ExecuteTemplate(w, "linkInterection", nil)
}
