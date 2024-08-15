package ascii

import (
	"fmt"
	"net/http"
)

// DownloadAscii handles the downloading of the ASCII art output
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, "405 : Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	out := r.FormValue("ascii") // Get the ascii value from the form
	if out == "" {
		ErrorHandler(w, "400 : Bad Request", http.StatusBadRequest)
		return
	}
	filename := "ascii.txt"
	// Set headers for the response
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(out))) // Set the content length

	// Write the output directly to the response
	w.Write([]byte(out))
}
