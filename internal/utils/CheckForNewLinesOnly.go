package utils

import "regexp"

func CheckForNewLinesOnly(regExpression string, inputWord string) (result string) {
	reg := regexp.MustCompile(regExpression)
	newWord := reg.ReplaceAllString(inputWord, "")
	if newWord == "" {
		for i := 0; i < len(inputWord)/2+1; i++ {
			result = result + "\n"
		}
		result = AddPreTagToString(result)
		return result
	}
	return ""
}
