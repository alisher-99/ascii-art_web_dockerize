package utils

import (
	"fmt"
	"net/http"
)

func CheckAsciiArtRequest(w http.ResponseWriter, r *http.Request) (status int, text string) {
	if r.Method == http.MethodGet || r.Method == http.MethodOptions || r.Method == http.MethodHead {
		return http.StatusNotFound, "404 Not found"
	}
	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, "405 Method Not Allowed"
	}
	if r.FormValue("styles") == "" {
		return http.StatusBadRequest, "400 Bad Request"
	}

	for _, char := range r.FormValue("inputText") {
		if (char < rune(32) || char > rune(126)) && (char != 13 && char != 10) {
			fmt.Println(char)
			return http.StatusBadRequest, "400 Bad Request"
		}
	}

	styleList := GetBannerList()
	inputBanner := r.FormValue("styles")
	if _, ok := styleList[inputBanner]; !ok {
		return http.StatusBadRequest, "400 Bad Request"
	}

	return 0, ""
}
