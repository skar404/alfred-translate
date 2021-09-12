package commands

import (
	"github.com/skar404/alfred-translate/global"
)

func getToken() {
	global.WF.NewItem("Set token").
		Subtitle("title").
		Arg(global.Flag.Value).
		Valid(true).
		Var("value", global.Flag.Value).
		Var("varname", Token)
}
