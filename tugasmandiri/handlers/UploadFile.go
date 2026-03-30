package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.MultipartForm.RemoveAll()

	uf, ufh, err := r.FormFile("media")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer uf.Close()

	flagStoragePath := "C:\\Users\\asus\\Documents\\StoreUpload\\"
	path := filepath.Join(flagStoragePath, ufh.Filename)

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f, err := os.Create(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := io.Copy(f, uf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("uploaded"))
}
