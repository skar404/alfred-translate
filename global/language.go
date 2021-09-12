package global

type L struct {
	ImgCode string
	Name    string
}

var langInfo = map[string]L{
	"ru": {"ru", "Russian"},
	"en": {"us", "English"},
}

func GetLangInfo(c string) L {
	return langInfo[c]
}
