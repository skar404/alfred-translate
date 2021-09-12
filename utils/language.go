package utils

import "github.com/skar404/alfred-translate/global"

func GetLangInfo(c string) global.L {
	return global.LangInfo[c]
}
