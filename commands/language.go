package commands

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/skar404/alfred-translate/global"
	"log"
	"strings"
)

func getConfigLanguage() map[string]bool {
	set := make(map[string]bool)

	languages := global.WF.Config.Get(global.Lang)
	if languages == "" {
		return set
	}
	splitLang := strings.Split(languages, ",")

	for _, l := range splitLang {
		set[l] = true
	}

	return set
}

func editLanguage(commandSub, value string) {
	log.Printf("languagesSet value commandSub")
	wf := global.WF
	lang := getConfigLanguage()

	switch commandSub {
	case global.Add:
		lang[value] = true
	case global.Delete:
		if _, ok := lang[value]; ok {
			delete(lang, value)
		}
	}

	var newLanguages []string
	for l := range lang {
		newLanguages = append(newLanguages, l)
	}

	if err := wf.Config.Set(global.Lang, strings.Join(newLanguages, ","), false).Do(); err != nil {
		wf.FatalError(err)
	}
}

func getLanguages() {
	wf := global.WF

	lang := getConfigLanguage()

	for _, l := range [][]string{
		{"ru", "Russian"},
		{"us", "English"}} {
		code := l[0]
		langName := l[1]

		name := fmt.Sprintf("➕ %s", langName)
		title := "Add new language"
		varName := fmt.Sprintf("%s:%s", global.Lang, global.Add)
		if _, ok := lang[code]; ok {
			name = fmt.Sprintf("➖ %s", langName)
			title = "Delete language"
			varName = fmt.Sprintf("%s:%s", global.Lang, global.Delete)
		}

		wf.NewItem(name).
			Subtitle(title).
			Icon(&aw.Icon{
				Value: fmt.Sprintf("./country-flags/img/%s.png", code),
				Type:  aw.IconTypeImage,
			}).
			Arg(code).
			Valid(true).
			Var("value", code).
			Var("varname", varName)
	}
}
