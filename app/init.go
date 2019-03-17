package app

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	EnvSetting.Token = os.Getenv("TOKEN")
	EnvSetting.FolderId = os.Getenv("FOLDER_ID")
	EnvSetting.WhiteListLang = strings.Split(os.Getenv("WHITE_LIST_LANG"), ",")
}
