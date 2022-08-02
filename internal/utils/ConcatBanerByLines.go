package utils

import (
	"errors"
	"strings"
)

func (args *InputArgs) ConcatBannerByLines(words []string, bannerMap map[rune]string) error {
	var lines [LinesHeight]string
	for i := 0; i < len(words); i++ {
		if len(words) > 1 && words[i] == "" {
			args.ConvertedString = args.ConvertedString + "\n"
			continue
		}

		for _, oneRune := range words[i] {
			oneMap := strings.Split(bannerMap[oneRune], "\n")

			if len(oneMap) == 1 && oneMap[0] == "" {
				return errors.New(ErrInvalidMapConvertString)
			}
			for j := 0; j < LinesHeight; j++ {
				lines[j] = lines[j] + oneMap[j]
			}
		}

		for j := 0; j < LinesHeight; j++ {
			if lines[j] > "" {
				args.ConvertedString = args.ConvertedString + lines[j] + "\n"
			}
			lines[j] = ""
		}
	}
	return nil
}
