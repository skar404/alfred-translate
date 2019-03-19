package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"alfred-translate/app"
	"alfred-translate/app/client"
)

func main() {
	app.Init()

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
		langList := app.GetTargetLanguageCode(app.DetectLang(opts.Text))

		for _, lang := range langList {
			data := client.TranslateText(opts.Text, lang)
			fmt.Printf("%+v\n", data)
		}
	case "auth":

	}
}
