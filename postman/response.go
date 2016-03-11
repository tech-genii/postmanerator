package postman

type Response struct {
	Status       string `json:"status"`
	ResponseCode struct {
		Code   int    `json:"code"`
		Name   string `json:"name"`
		Detail string `json:"detail"`
	} `json:"responseCode"`
	Time    interface{} `json:"time"`
	Headers []struct {
		Name        string `json:"name"`
		Key         string `json:"key"`
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"headers"`
	Cookies     []interface{} `json:"cookies"`
	Mime        string        `json:"mime"`
	Text        string        `json:"text"`
	Language    string        `json:"language"`
	RawDataType string        `json:"rawDataType"`
	State       struct {
		Size string `json:"size"`
	} `json:"state"`
	PreviewType            string      `json:"previewType"`
	SearchResultScrolledTo interface{} `json:"searchResultScrolledTo"`
	ForceNoPretty          bool        `json:"forceNoPretty"`
	Write                  bool        `json:"write"`
	Empty                  bool        `json:"empty"`
	Failed                 bool        `json:"failed"`
	IsSample               bool        `json:"isSample"`
	ScrollToResult         bool        `json:"scrollToResult"`
	RunTests               bool        `json:"runTests"`
	ID                     string      `json:"id"`
	Name                   string      `json:"name"`
	Request                struct {
		URL     string `json:"url"`
		Headers []struct {
			Key     string `json:"key"`
			Value   string `json:"value"`
			Name    string `json:"name"`
			Enabled bool   `json:"enabled"`
		} `json:"headers"`
		Data     string `json:"data"`
		Method   string `json:"method"`
		DataMode string `json:"dataMode"`
	} `json:"request"`
}
