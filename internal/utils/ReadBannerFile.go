package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func (banner *BannerMaps) ReadBannerFile() error {
	file, err := os.Open(banner.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	symbol := ""
	count := 0
	oneRune := FirstRuneInMap
	banner.BannerMap = make(map[rune]string, MapSize)
	for fileScanner.Scan() {
		symbol = symbol + "\n" + fileScanner.Text()
		count++
		if count%BannersHeight == 1 && fileScanner.Text() != "" {
			fmt.Println(fileScanner.Text())
			return errors.New(ErrInvalidBannerFile)
		}

		if count%BannersHeight == 0 {
			banner.BannerMap[oneRune] = symbol[2:]
			symbol = ""
			oneRune++

		}
	}

	if count%BannersHeight != 0 {
		return errors.New(ErrInvalidBannerFile)
	}

	if err := fileScanner.Err(); err != nil {
		return err
	}

	return nil
}
