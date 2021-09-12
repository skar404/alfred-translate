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
	case global.Lang:
		editLanguage(commandSub, value)
	case global.Back:

	default:
		if err := wf.Config.Set(key, value, false).Do(); err != nil {
			wf.FatalError(err)
		}
	}

	if err := wf.Alfred.RunTrigger(global.NameAlfredExtensionSetting, ""); err != nil {
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

	back := true

	switch key {
	case global.Lang:
		getLanguages()
	case global.Token:
		getToken()
	default:
		back = false
		wf.NewItem("Add/delete language").
			Subtitle("↩ to edit").
			Valid(true).
			Var("name", global.Lang)
		wf.NewItem("Set token").
			Subtitle("↩ to edit").
			Valid(true).
			Var("name", global.Token)
	}

	if back {
		wf.NewItem("↩ back").
			Valid(true).
			Var("varname", global.Back)
	}
}
