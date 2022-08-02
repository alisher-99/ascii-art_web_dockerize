package utils

type InputArgs struct {
	InputWord       string
	ConvertedString string
}

type BannerMaps struct {
	BannerMap map[rune]string
	FileName  string
}

const (
	ErrInvalidBannerFile        string = "invalid data in the banner file"
	ErrInvalidMapConvertString  string = "invalid data in the banner map"
	ErrInvalidCharConvertString string = "invalid input char"
	BannerFileName              string = "./internal/files/standard.txt"

	MapSize        int  = 95
	FirstRuneInMap rune = ' '
	BannersHeight  int  = 9
	LinesHeight    int  = 8
)

type ViewData struct { //
	Title      string
	StylesList string // fonts
	Banner     string // result
}

func GetBannerList() map[string]string {
	bannersList := map[string]string{
		"standard":   "<option value=\"standard\">Standard</option>",
		"shadow":     "<option value=\"shadow\">Shadow</option>",
		"thinkertoy": "<option value=\"thinkertoy\">Thinkertoy</option>",
	}

	return bannersList
}
