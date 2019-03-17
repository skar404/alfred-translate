package main

import (
	"alfred-translate/app"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var opts app.Opts
	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	switch opts.Command {
	case "translate":
		app.InitEnv()
		lang := app.DetectLang(opts.Text)

		data := app.TranslateText(opts.Text, lang)

		fmt.Printf("%+v\n", data)
	case "auth":

	}
}
