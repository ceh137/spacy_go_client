package spacy_go_client

import (
	"encoding/json"
	"net/http"
)

type ModelNotExistsError struct {
	ModelNumber int
}

func (err ModelNotExistsError) Error() string {
	str := `
This model is not available in SpaCy.
Available models and constants to use:
en EnModel
de DeModel
es EsModel
fr FrModel
pt PtModel
it ItModel
nl NlModel
`
	return str
}

func getDataFromJson[T any](response *http.Response) (T, error) {
	var result T
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func convertToJson(data any) ([]byte, error) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}
	return jsonString, err
}

type GetModelsResponse = []string

func (sc *SpacyClient) GetModels() (GetModelsResponse, error) {
	url := sc.Url + "/models"
	response, err := http.Get(url)
	if err != nil {
		return GetModelsResponse{}, err
	}
	respBody, err := getDataFromJson[GetModelsResponse](response)
	if err != nil {
		return GetModelsResponse{}, err
	}

	return respBody, nil
}

//type GetModelSchemaResponse struct {
//	DepTypes []string `json:"dep_types"`
//	EntTypes []string `json:"ent_types"`
//	PosTypes []string `json:"pos_types"`
//}
//
//func (sc *SpacyClient) GetModelSchema(model string) (GetModelSchemaResponse, error) {
//	url := fmt.Sprintf("%s/%s/schema", sc.Url, model)
//	response, err := http.Get(url)
//	if err != nil {
//		return GetModelSchemaResponse{}, err
//	}
//	respBody, err := getDataFromJson[GetModelSchemaResponse](response)
//	if err != nil {
//		return GetModelSchemaResponse{}, err
//	}
//
//	return respBody, nil
//}

type VersionResponse struct {
	Spacy string `json:"spacy"`
}

func (sc *SpacyClient) GetVersion() (VersionResponse, error) {
	url := sc.Url + "/version"
	response, err := http.Get(url)
	if err != nil {
		return VersionResponse{}, err
	}
	respBody, err := getDataFromJson[VersionResponse](response)
	if err != nil {
		return VersionResponse{}, err
	}

	return respBody, nil
}
