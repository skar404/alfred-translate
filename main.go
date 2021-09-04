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
	token := wf.Config.Get("api_token")

	text := translateText("hello world", token)

	// Add a "Script Filter" result
	wf.NewItem(text)
	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
