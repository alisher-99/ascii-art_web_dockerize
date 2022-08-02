package utils

import "net/http"

func CheckIndexRequest(w http.ResponseWriter, r *http.Request) (status int, text string) {
	// валидация пути
	if r.URL.Path != "/" {
		return http.StatusNotFound, "404 Not Found"
	}

	// валидация метода
	if r.Method == http.MethodOptions || r.Method == http.MethodHead {
		return http.StatusNotFound, "404 Not found"
	}
	if r.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, "405 Method Not Allowed"
	}

	return 0, ""
}
