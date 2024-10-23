package report_helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ServeReport serves the pretty report via a web UI.
func ServeReport() {
	http.HandleFunc("/report", reportHandler)
	fmt.Println("Serving report at http://localhost:8080/report")
	http.ListenAndServe(":8080", nil)
}

// reportHandler reads and displays the report content.
func reportHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./reports/pretty-report.txt")
	if err != nil {
		fmt.Fprintf(w, "Error reading report file: %v", err)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}
