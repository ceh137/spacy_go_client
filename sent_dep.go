package spacy_go_client

import (
	"net/http"
	"strings"
)

type sentDepRequestBody struct {
	Text  string `json:"text"`
	Model string `json:"model"`
}

type sentDepItem struct {
	Sentence string          `json:"sentence"`
	DepParse DepResponseType `json:"dep_parse"`
}

type SentDepResponseType = []sentDepItem

func (sc *SpacyClient) GetSentDeps(text string) (SentDepResponseType, error) {
	url := sc.Url + "/sents_dep"
	jsonToSend, err := convertToJson(sentDepRequestBody{
		Text:  text,
		Model: sc.Model,
	})
	if err != nil {
		return SentDepResponseType{}, err
	}
	reader := strings.NewReader(string(jsonToSend))
	response, err := http.Post(url, "application/json", reader)
	if err != nil {
		return SentDepResponseType{}, err
	}
	respBody, err := getDataFromJson[SentDepResponseType](response)
	if err != nil {
		return SentDepResponseType{}, err
	}

	return respBody, nil
}
