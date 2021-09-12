package commands

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/skar404/alfred-translate/global"
	yandex_translate "github.com/skar404/alfred-translate/yandex-translate"
	"log"
)

func Translate() {
	wf := global.WF

	for lang := range getConfigLanguage() {
		log.Printf("Translate %s", lang)
		info := global.GetLangInfo(lang)

		text, err := yandex_translate.TranslateText(global.Flag.Value, lang)
		if err != nil {
			wf.NewItem(fmt.Sprintf("Error %s", err))
			return
		}

		wf.NewItem(text).Icon(&aw.Icon{
			Value: fmt.Sprintf("./country-flags/img/%s.png", info.ImgCode),
			Type:  aw.IconTypeImage,
		})
	}
}
