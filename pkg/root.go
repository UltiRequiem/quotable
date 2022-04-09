package quotable

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Quotable struct {
	Id           string   `json:"_id"`
	Tags         []string `json:"tags"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}

const (
	quotableApiUrl = "https://api.quotable.io/random"
)

func GetQuotable() (Quotable, error) {
	response, getError := http.Get(quotableApiUrl)

	if getError != nil {
		return Quotable{}, getError
	}

	responseData, readError := ioutil.ReadAll(response.Body)

	if readError != nil {
		return Quotable{}, readError
	}

	var parsedData Quotable

	json.Unmarshal(responseData, &parsedData)

	return parsedData, nil
}
