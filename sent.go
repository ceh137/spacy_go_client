package spacy_go_client

import (
	"net/http"
	"strings"
)

type sentRequestBody struct {
	Text  string `json:"text"`
	Model string `json:"model"`
}
type SentResponseType = []string

func (sc *SpacyClient) GetSents(text string) (SentResponseType, error) {
	url := sc.Url + "/sents"
	jsonToSend, err := convertToJson(sentRequestBody{
		Text:  text,
		Model: sc.Model,
	})
	if err != nil {
		return SentResponseType{}, err
	}
	reader := strings.NewReader(string(jsonToSend))
	response, err := http.Post(url, "application/json", reader)
	if err != nil {
		return SentResponseType{}, err
	}
	respBody, err := getDataFromJson[SentResponseType](response)
	if err != nil {
		return SentResponseType{}, err
	}

	return respBody, nil
}
