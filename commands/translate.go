package commands

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/skar404/alfred-translate/global"
	"github.com/skar404/alfred-translate/utils"
	yandex_translate "github.com/skar404/alfred-translate/yandex-translate"
	"log"
)

func Translate() {
	wf := global.WF

	if global.Flag.Value == "" {
		wf.NewItem("Enter translate text ... ")
		return
	}

	for lang := range getConfigLanguage() {
		log.Printf("Translate %s", lang)
		info := utils.GetLangInfo(lang)

		text, err := yandex_translate.TranslateText(global.Flag.Value, lang)
		if err != nil {
			wf.NewItem(fmt.Sprintf("Error translate, err: %s", err)).
				Subtitle(fmt.Sprintf("Full err: %s", err)).Valid(false)
			return
		}

		wf.NewItem(text).Icon(&aw.Icon{
			Value: fmt.Sprintf("./country-flags/img/%s.png", info.ImgCode),
			Type:  aw.IconTypeImage,
		}).Valid(true).Var("text", text)
	}
}
