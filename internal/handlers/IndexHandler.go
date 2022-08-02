package handlers

import (
	"asciiWeb/internal/utils"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// валидация запроса
	if status, text := utils.CheckIndexRequest(w, r); status != 0 {
		http.Error(w, text, status)
		return
	}

	// формирование шаблона к выводу и проверка ошибки
	tmpl, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	// формирование данных для шаблона
	data := new(utils.ViewData)
	data.FormHtmlTempData("")
	// исполнение шаблона
	tmpl.Execute(w, data)
}
