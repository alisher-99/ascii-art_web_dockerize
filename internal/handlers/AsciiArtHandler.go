package handlers

import (
	"asciiWeb/internal/utils"
	"net/http"
	"text/template"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// валидация метода и тела запроса
	if status, text := utils.CheckAsciiArtRequest(w, r); status != 0 {
		http.Error(w, text, status)
		return
	}

	args := new(utils.InputArgs)
	args.InputWord = r.FormValue("inputText")
	banner := new(utils.BannerMaps)
	banner.FileName = "./internal/files/" + r.FormValue("styles") + ".txt"

	// чтение файла с баннерами и проверка ошибки
	if err := banner.ReadBannerFile(); err != nil {
		http.Error(w, "500 Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// конвертация текста по баннеру и проверка ошибки
	if err := args.ConvertString(banner.BannerMap); err != nil {
		http.Error(w, "500 Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// формирование данных для шаблона
	data := new(utils.ViewData)
	data.FormHtmlTempData(args.ConvertedString)

	// формирование шаблона к выводу и проверка ошибки
	tmpl, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// исполнение шаблона
	tmpl.Execute(w, data)
}
