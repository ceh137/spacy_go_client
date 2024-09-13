package spacy_go_client

import (
	"net/http"
	"strings"
)

type entRequestBody struct {
	Text  string `json:"text"`
	Model string `json:"model"`
}

type EntItem struct {
	End   int    `json:"end"`
	Start int    `json:"start"`
	Type  string `json:"type"`
	Text  string `json:"text"`
}

type EntResponseType = []EntItem

func (sc *SpacyClient) GetEnts(text string) (EntResponseType, error) {
	url := sc.Url + "/ent"
	jsonToSend, err := convertToJson(entRequestBody{
		Text:  text,
		Model: sc.Model,
	})
	if err != nil {
		return EntResponseType{}, err
	}
	reader := strings.NewReader(string(jsonToSend))
	response, err := http.Post(url, "application/json", reader)
	if err != nil {
		return EntResponseType{}, err
	}
	respBody, err := getDataFromJson[EntResponseType](response)
	if err != nil {
		return EntResponseType{}, err
	}

	return respBody, nil
}
