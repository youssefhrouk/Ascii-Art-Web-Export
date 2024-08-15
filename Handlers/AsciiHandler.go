package ascii

import (
	"html/template"
	"net/http"

	ascii "ascii/functions"
)

// handles the ASCII art generation requests
func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, "405 : Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	input := r.FormValue("input")
	banner := r.FormValue("banner")
	if input == "" || banner == "" {
		ErrorHandler(w, "400 : Bad Request - Input and banner selection are required", http.StatusBadRequest)
		return
	}

	if len(input) > 1000 {
		ErrorHandler(w, "400 : Bad Request - Input exceeds the maximum allowed length of 1000 characters", http.StatusBadRequest)
		return
	}

	output, status := ascii.PrintAndSplit(input, banner)
	if status != http.StatusOK {
		ErrorHandler(w, output, status)
		return
	}

	tmpl, err := template.ParseFiles("templates/ascii-art.html")
	if err != nil {
		ErrorHandler(w, "500 : Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Message: output,
	}

	if err := tmpl.Execute(w, data); err != nil {
		ErrorHandler(w, "500 : Internal Server Error", http.StatusInternalServerError)
	}
}
