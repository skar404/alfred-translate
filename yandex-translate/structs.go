package yandex_translate

type ResponseDone struct {
	Translations []struct {
		Text                 string `json:"text"`
		DetectedLanguageCode string `json:"detectedLanguageCode"`
	} `json:"translations"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details []struct {
		Type      string `json:"@type"`
		RequestId string `json:"requestId"`
	} `json:"details"`
}
