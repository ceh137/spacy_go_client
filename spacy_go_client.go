package spacy_go_client

const (
	EnModel = "en"
	DeModel = "de"
	EsModel = "es"
	FrModel = "fr"
	PtModel = "pt"
	ItModel = "it"
	NlModel = "nl"
)

type SpacyClient struct {
	Url   string
	Model string
}
