package spacy_go_client

import (
	"net/http"
	"strings"
)

const (
	RightArrow = iota
	LeftArrow
)

type Arc struct {
	Dir   string
	Start int
	End   int
	Label string
	Text  string
}

type Word struct {
	Tag  string
	Text string
}

type depRequestBody struct {
	Text                string `json:"text"`
	Model               string `json:"model"`
	CollapsePunctuation bool   `json:"collapse_punctuation"`
	CollapsePhrases     bool   `json:"collapse_phrases"`
}

type DepResponseType struct {
	Arcs  []Arc  `json:"arcs"`
	Words []Word `json:"words"`
}

func (sc *SpacyClient) GetDeps(text string, collapsePunctuation bool, collapsePhrases bool) (DepResponseType, error) {
	url := sc.Url + "/dep"
	jsonToSend, err := convertToJson(depRequestBody{
		Text:                text,
		Model:               sc.Model,
		CollapsePunctuation: collapsePunctuation,
		CollapsePhrases:     collapsePhrases,
	})
	if err != nil {
		return DepResponseType{}, err
	}
	reader := strings.NewReader(string(jsonToSend))
	response, err := http.Post(url, "application/json", reader)
	if err != nil {
		return DepResponseType{}, err
	}
	respBody, err := getDataFromJson[DepResponseType](response)
	if err != nil {
		return DepResponseType{}, err
	}

	return respBody, nil
}
