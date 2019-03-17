package app

type Opts struct {
	// Example of automatic marshalling to desired type (uint)
	Text    string `long:"text" description:"Translate text"`
	Type    string `long:"type" choice:"alfred" choice:"cli"`
	Command string `long:"command" choice:"auth" choice:"translate"`
}

type Env struct {
	Token         string
	FolderId      string
	WhiteListLang []string
}

var EnvSetting Env
