package utils

func (data *ViewData) FormHtmlTempData(banner string) {
	data.Title = "ASCII ART WEB"

	styleList := GetBannerList()
	for _, value := range styleList {
		data.StylesList = data.StylesList + value
	}
	data.Banner = banner
}
