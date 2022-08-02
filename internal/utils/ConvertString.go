package utils

func (args *InputArgs) ConvertString(bannerMap map[rune]string) error {
	if tempResult := CheckForNewLinesOnly(`^(\\n)*$`, args.InputWord); tempResult != "" {
		args.ConvertedString = AddPreTagToString(tempResult)
		return nil
	}
	if tempResult := CheckForNewLinesOnly(`^(\r\n)*$`, args.InputWord); tempResult != "" {
		args.ConvertedString = AddPreTagToString(tempResult)
		return nil
	}

	words := SplitInputWord(args.InputWord)
	if err := args.ConcatBannerByLines(words, bannerMap); err != nil {
		return err
	}

	args.ConvertedString = AddPreTagToString(args.ConvertedString)
	return nil
}
