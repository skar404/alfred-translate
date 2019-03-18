package app

import (
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

func Init() {
	ProjectDir()
	ProjectEnv()
}

func ProjectDir() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func ProjectEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	EnvSetting.Token = os.Getenv("TOKEN")
	EnvSetting.FolderId = os.Getenv("FOLDER_ID")
	EnvSetting.WhiteListLang = strings.Split(os.Getenv("WHITE_LIST_LANG"), ",")
}
