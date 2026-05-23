package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("C:/Dev/sprint6-final/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Cotent-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Check method (POST)
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Gets the file from the form
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer file.Close()
	// Read the file content
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	filename := "result" + time.Now().Format("20060102_150405") + filepath.Ext(header.Filename)

	err = os.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(result))
}
