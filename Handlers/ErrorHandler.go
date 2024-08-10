package ascii

import (
	"html/template"
	"net/http"
)

type ErrorData struct {
	Message string
}

func ErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Prepare the data for the template
	data := ErrorData{
		Message: message,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
