package app

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)


type Student struct {
	FolderId    string `json:"folder_id"`
	Texts []string `json:"texts"`
	TargetLanguageCode string `json:"targetLanguageCode"`
}


func TranslateText(text string) {
	body := &Student{
		FolderId:    EnvSetting.FolderId,
		Texts: []string{text},
		TargetLanguageCode: "en",
	}

	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", "https://translate.api.cloud.yandex.net/translate/v2/translate", buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + EnvSetting.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	println(bodyString)
}
