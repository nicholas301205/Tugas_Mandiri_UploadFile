package routes

import (
	"net/http"
	"tugasmandiri/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/uploadfile", handlers.UploadFile)

	http.ListenAndServe(":8081", nil)
}
