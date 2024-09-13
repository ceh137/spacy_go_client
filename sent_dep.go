package spacy_go_client

import (
	"net/http"
	"strings"
)

type sendDepRequestBody struct {
	Text  string `json:"text"`
	Model string `json:"model"`
}

type sendDepItem struct {
	Sentence string          `json:"sentence"`
	DepParse DepResponseType `json:"dep_parse"`
}

type SendDepResponseType = []sendDepItem

func (sc *SpacyClient) GetSentDeps(text string) (SendDepResponseType, error) {
	url := sc.Url + "/sents_dep"
	jsonToSend, err := convertToJson(sendDepRequestBody{
		Text:  text,
		Model: sc.Model,
	})
	if err != nil {
		return SendDepResponseType{}, err
	}
	reader := strings.NewReader(string(jsonToSend))
	response, err := http.Post(url, "application/json", reader)
	if err != nil {
		return SendDepResponseType{}, err
	}
	respBody, err := getDataFromJson[SendDepResponseType](response)
	if err != nil {
		return SendDepResponseType{}, err
	}

	return respBody, nil
}
