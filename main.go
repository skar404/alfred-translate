package main

// Package is called aw
import (
	"fmt"
	"github.com/deanishe/awgo"
	"github.com/go-resty/resty/v2"
	"strings"
)

// Workflow is the main API
var wf *aw.Workflow

type TranslateBody struct {
	FolderId           string   `json:"folderId"`
	Texts              []string `json:"texts"`
	TargetLanguageCode string   `json:"targetLanguageCode"`
}

type TranslateSuccess struct {
	Translations []struct {
		Text                 string `json:"text"`
		DetectedLanguageCode string `json:"detectedLanguageCode"`
	} `json:"translations"`
}

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

func translateText(text, token string) string {
	client := resty.New()
	resp, _ := client.R().
		EnableTrace().
		SetBody(TranslateBody{
			Texts:              strings.Split(text, " "),
			TargetLanguageCode: "ru",
		}).
		SetHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Api-Key %s", token),
			"Content-Type":  "application/json",
		}).
		SetResult(&TranslateSuccess{}).
		Post("https://translate.api.cloud.yandex.net/translate/v2/translate")

	res := resp.Result().(*TranslateSuccess)

	textRaw := make([]string, len(res.Translations))

	for i, element := range res.Translations {
		textRaw[i] = element.Text
	}

	return strings.Join(textRaw, " ")
}

// Your workflow starts here
func run() {
	args := wf.Args()

	if len(args) < 2 {
		wf.WarnEmpty("No matching folders found", "Try a different query?")
		wf.SendFeedback()
		return
	}

	command := args[0]

	switch command {
	case "translate":
		token := wf.Config.Get("api_token")
		text := strings.Join(args[1:], " ")
		reqTest := translateText(text, token)
		wf.NewItem(reqTest).
			Var("text", reqTest).
			Arg("ok").
			Valid(true)
	case "set_token":
		wf.NewItem("Set yandex API token").
			Subtitle("copy api token in yandex cloud and input ").
			Arg(args[1]).
			ActionForType("file", "copy api token in yandex cloud and input").Valid(true)
	case "set_token_done":
		token := args[1]
		_ = wf.Config.Set("api_token", token, false).Do()
	}

	wf.WarnEmpty("No matching folders found", "Try a different query?")
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
