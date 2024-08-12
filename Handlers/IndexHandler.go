package ascii

import (
	ascii "ascii/functions"
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

func DownloadAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
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

	filename := "ascii-art.txt"
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(output))

}
