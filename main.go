package main

import (
	"flag"
	aw "github.com/deanishe/awgo"
	"github.com/skar404/alfred-translate/commands"
	"github.com/skar404/alfred-translate/global"
	"log"
)

func init() {
	global.WF = aw.New()
	flag.StringVar(&global.Flag.Command, "command", "", "command type")
	flag.StringVar(&global.Flag.Value, "value", "", "")
	flag.StringVar(&global.Flag.SetKey, "set", "", "save a value")
	flag.StringVar(&global.Flag.GetKey, "get", "", "enter a new value")
}

func run() {
	flag.Parse()
	log.Printf("flag flag=%+v args=%+v", global.Flag, flag.Args())

	command := global.Flag.Command

	switch command {
	case global.Setting:
		commands.Setting()
	case global.Translate:

	}

	global.WF.WarnEmpty("No Matching Items", "Try a different query?")
	global.WF.SendFeedback()
}

func main() {
	global.WF.Run(run)
}
