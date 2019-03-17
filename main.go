package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"

	"alfred_translate/app"
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
		initEnv()
		app.TranslateText(opts.Text)
	case "auth":
		fmt.Printf("Go to link: sss ")
	}
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		os.Exit(0)
	}

	app.EnvSetting.Token = os.Getenv("TOKEN")
	app.EnvSetting.FolderId = os.Getenv("FOLDER_ID")
}
