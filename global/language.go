package global

type L struct {
	ImgCode string
	Name    string
}

var LangInfo = map[string]L{
	"ru": {"ru", "Russian"},
	"en": {"us", "English"},
}
