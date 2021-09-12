package commands

import (
	"fmt"
	"github.com/skar404/alfred-translate/global"
	yandex_translate "github.com/skar404/alfred-translate/yandex-translate"
)

func getToken() {
	wf := global.WF
	apiToken := global.Flag.Value

	if apiToken == "" {
		wf.NewItem(fmt.Sprintf("Enter new %s", global.Token))
		return
	}

	valid := yandex_translate.TestRequest(apiToken)

	if valid {
		wf.NewItem("Set token").
			Subtitle("Token is valid").
			Arg(apiToken).
			Valid(true).
			Var("value", apiToken).
			Var("varname", global.Token)
		return
	}

	wf.NewItem("Not valid token").
		Subtitle("Not valid token, read readme")
}
