package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"alfred-translate/app"
)

type RequestData struct {
	FolderId           string   `json:"folder_id"`
	Texts              []string `json:"texts"`
	TargetLanguageCode string   `json:"targetLanguageCode"`
}

type ReqTranslateText struct {
	Text                 string `json:"text"`
	DetectedLanguageCode string `json:"detectedLanguageCode"`
}

type ReqErrorDetails struct {
	Type string `json:"@type"`
	RequestId string `json:"requestId"`
}

type ReqTranslate struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Details []ReqErrorDetails `json:"details"`

	Translations []ReqTranslateText `json:"translations"`
}


func jsonHttpClient(method string, url string, body, target interface{}) {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+app.EnvSetting.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		panic(err)
	}
}

func TranslateText(text string, targetLanguage string) ReqTranslate {
	body := &RequestData{
		FolderId:           app.EnvSetting.FolderId,
		Texts:              []string{text},
		TargetLanguageCode: targetLanguage,
	}
	var reqData ReqTranslate

	jsonHttpClient("POST", "https://translate.api.cloud.yandex.net/translate/v2/translate", body, &reqData)
	return reqData
}
