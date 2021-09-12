package yandex_translate

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/skar404/alfred-translate/global"
	"net/http"
	"strings"
)

const (
	Url = "https://translate.api.cloud.yandex.net/translate/v2/"
)

var BadRequest = errors.New("bad request")

func TestRequest(token string) bool {
	client := resty.New()

	res, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Api-Key %s", token)).
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("%slanguages", Url))

	if err != nil {
		return false
	}

	return res.StatusCode() == http.StatusOK
}

func TranslateText(text, lang string) (string, error) {
	token := global.WF.Config.Get(global.Token)
	client := resty.New()

	doneRes := ResponseDone{}
	errRes := ResponseError{}

	_, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Api-Key %s", token)).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"texts":              strings.Split(text, " "),
			"targetLanguageCode": lang,
		}).
		SetResult(&doneRes).
		SetError(&errRes).
		Post(fmt.Sprintf("%translate", Url))

	if err != nil {
		return "", fmt.Errorf(errRes.Message, BadRequest)
	}

	var translateText []string

	for _, v := range doneRes.Translations {
		translateText = append(translateText, v.Text)
	}

	return strings.Join(translateText, " "), nil

}
