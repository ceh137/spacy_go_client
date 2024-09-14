
## Go client For SpaCy Server
A basic Client for SpaCy API Server implemented in Go

Notes:

- Was tested with docker image of the server: [jgontrum/spacyapi:en_v2](https://github.com/jgontrum/spacy-api-docker)

### Installation:
Run Go get to download this package:
```bash
go get github.com/ceh137/spacy_go_client 
```
Import package into files you would like to use in:
```go
package main

import (
	spy "github.com/ceh137/spacy_go_client"
)
```

### Usage
Two simple steps:

- Create `SpacyClient`.
- Use it to call methods

```go
func main() {
	sc := spy.SpacyClient{
		Model: spy.EnModel,
		Url:   "localhost:8080",
	}
	
	sentDeps, _ := sc.GetSentDeps("Some text")
	ents, _ := sc.GetEnts("Some text")
	deps, _ := sc.GetDeps("Some text", true, true)
}
```

### Available Methods And Constants
To understand the terminology better please refer to the original [SpaCy Documentation](https://spacy.io/usage/linguistic-features)

Feel free to use these constants to set the model in your client:
```go
const (
	EnModel = "en"
	DeModel = "de"
	EsModel = "es"
	FrModel = "fr"
	PtModel = "pt"
	ItModel = "it"
	NlModel = "nl"
)
```

### `sc.GetSentDeps(text string)`
Returns Sentences and Dependency parses

---

### `sc.GetSents(text string)`
Returns Sentences from text

---

### `sc.GetEnts(text string)`
Returns Entities found in text

---

### `sc.GetDeps(text string)`
Returns Dependencies found in text

---

### `sc.GetVersion()`
Returns the version od spacy on your server

---

### `sc.GetModels()`
Returns available on your server models
