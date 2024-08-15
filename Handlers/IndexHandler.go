package ascii

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Message string
}

// handles requests to the root URL
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, "405 : Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	if r.URL.Path != "/" {
		ErrorHandler(w, "404 : Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "500 : Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		ErrorHandler(w, "500 : Internal Server Error", http.StatusInternalServerError)
	}
}
