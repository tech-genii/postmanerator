package postman

import "strings"

type Request struct {
	ID               string        `json:"id"`
	RawHeaders       string        `json:"headers"`
	URL              string        `json:"url"`
	PreRequestScript string        `json:"preRequestScript"`
	PathVariables    interface{}   `json:"pathVariables"`
	Method           string        `json:"method"`
	Data             []interface{} `json:"data"`
	DataMode         string        `json:"dataMode"`
	Version          int           `json:"version"`
	Tests            string        `json:"tests"`
	CurrentHelper    string        `json:"currentHelper"`
	HelperAttributes interface{}   `json:"helperAttributes"`
	Time             interface{}   `json:"time"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	CollectionID     string        `json:"collectionId"`
	Responses        []Response    `json:"responses"`
	RawModeData      string        `json:"rawModeData"`
}

type Header struct {
	Name  string
	Value string
}

func (req Request) Headers() (headers []Header) {
	rawHeadersKeyValList := strings.Split(req.RawHeaders, "\n")
	for _, rawHeaderKeyVal := range rawHeadersKeyValList {
		keyVal := strings.Split(rawHeaderKeyVal, ": ")
		if len(keyVal) != 2 {
			continue
		}
		headers = append(headers, Header{keyVal[0], keyVal[1]})
	}
	return
}
