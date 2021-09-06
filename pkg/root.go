// Exports Quotable Struct and FetchQuotabler Method
package quotable

import "encoding/json"

// Quotable API Response Schema
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

func GetQuotableObject() Quotable {
	var data string = parseHttpResponse(fetch("https://api.quotable.io/random"))

	var parsedData Quotable
	json.Unmarshal([]byte(data), &parsedData)

	return parsedData
}

func GetRandomQuoteAndAuthor() (string, string) {
	quotableFetch := GetQuotableObject()
	return quotableFetch.Content, quotableFetch.Author
}
