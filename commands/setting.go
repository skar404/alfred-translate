package commands

import (
	aw "github.com/deanishe/awgo"
	"github.com/skar404/alfred-translate/global"
	"log"
	"strings"
)

func runSet(key, value string) {
	wf := global.WF

	wf.Configure(aw.TextErrors(true))

	log.Printf("saving %#v to %s ...", value, key)

	commandList, commandSub := strings.Split(key, ":"), ""
	cLen := len(commandList)
	if cLen < 1 {
		return
	} else if cLen > 1 {
		commandSub = commandList[1]
	}
	command := commandList[0]

	switch command {
	case Lang:
		editLanguage(commandSub, value)
	default:
		if err := wf.Config.Set(key, value, false).Do(); err != nil {
			wf.FatalError(err)
		}
	}

	if err := wf.Alfred.RunTrigger(NameAlfredExtensionSetting, ""); err != nil {
		wf.FatalError(err)
	}

	log.Printf("saved %#v to %s", value, key)
}

func Setting() {
	wf := global.WF
	key := global.Flag.GetKey

	if global.Flag.SetKey != "" {
		runSet(global.Flag.SetKey, global.Flag.Value)
		return
	}

	switch key {
	case Lang:
		getLanguages()
	case Token:
		getToken()
	default:
		wf.NewItem("Add/delete language").
			Subtitle("↩ to edit").
			Valid(true).
			Var("name", Lang)
		wf.NewItem("Set token").
			Subtitle("↩ to edit").
			Valid(true).
			Var("name", Token)
	}
}
