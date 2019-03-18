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
		//lang := app.DetectLang(opts.Text)

		data := client.TranslateText(opts.Text, "en")

		fmt.Printf("%+v\n", data)
	case "auth":

	}
}
