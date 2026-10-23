package quotable

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	// DefaultAPIURL is the public Quotable host, used when nothing else is configured.
	DefaultAPIURL = "https://api.quotable.io"

	// APIURLEnv is the environment variable that overrides DefaultAPIURL.
	APIURLEnv = "QUOTABLE_API_URL"
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

// BaseURL works out which server to talk to: the explicit value, then the
// QUOTABLE_API_URL environment variable, then DefaultAPIURL. Trailing slashes
// are removed, so "https://host" and "https://host/" are the same.
func BaseURL(explicit string) string {
	url := explicit

	if url == "" {
		url = os.Getenv(APIURLEnv)
	}

	if url == "" {
		url = DefaultAPIURL
	}

	return strings.TrimRight(url, "/")
}

// GetQuotable returns a random quote from the configured server
// (QUOTABLE_API_URL, or DefaultAPIURL when it is not set).
func GetQuotable() (Quotable, error) {
	return GetQuotableFrom("")
}

// GetQuotableFrom returns a random quote from the Quotable server at baseURL,
// for example "https://quotes.example.com". An empty baseURL falls back to
// QUOTABLE_API_URL and then to DefaultAPIURL.
func GetQuotableFrom(baseURL string) (Quotable, error) {
	url := BaseURL(baseURL) + "/random"

	response, err := http.Get(url)

	if err != nil {
		return Quotable{}, err
	}

	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return Quotable{}, fmt.Errorf("quotable: %s answered with status %d", url, response.StatusCode)
	}

	var quote Quotable

	if err := json.NewDecoder(response.Body).Decode(&quote); err != nil {
		return Quotable{}, fmt.Errorf("quotable: could not read the answer of %s: %w", url, err)
	}

	return quote, nil
}
