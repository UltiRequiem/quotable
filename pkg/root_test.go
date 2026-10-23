package quotable

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func withEnv(t *testing.T, value string) {
	t.Helper()

	previous, had := os.LookupEnv(APIURLEnv)

	if value == "" {
		os.Unsetenv(APIURLEnv)
	} else {
		os.Setenv(APIURLEnv, value)
	}

	t.Cleanup(func() {
		if had {
			os.Setenv(APIURLEnv, previous)
		} else {
			os.Unsetenv(APIURLEnv)
		}
	})
}

func TestBaseURLPrecedence(t *testing.T) {
	withEnv(t, "")

	if got := BaseURL(""); got != DefaultAPIURL {
		t.Fatalf("default: got %q", got)
	}

	withEnv(t, "https://env.example.com/")

	if got := BaseURL(""); got != "https://env.example.com" {
		t.Fatalf("env var: got %q", got)
	}

	if got := BaseURL("https://explicit.example.com///"); got != "https://explicit.example.com" {
		t.Fatalf("explicit value must win and lose its trailing slashes: got %q", got)
	}
}

func newServer(t *testing.T, status int, body string) (*httptest.Server, *string) {
	t.Helper()

	var path string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))

	t.Cleanup(server.Close)

	return server, &path
}

func TestGetQuotableFromUsesTheGivenServer(t *testing.T) {
	server, path := newServer(t, 200, `{"_id":"1","content":"Be yourself.","author":"Someone","length":12}`)

	quote, err := GetQuotableFrom(server.URL + "/api/")

	if err != nil {
		t.Fatal(err)
	}

	if quote.Content != "Be yourself." || quote.Author != "Someone" {
		t.Fatalf("unexpected quote: %+v", quote)
	}

	if *path != "/api/random" {
		t.Fatalf("expected /api/random, the server saw %q", *path)
	}
}

func TestGetQuotableReadsTheEnvironment(t *testing.T) {
	server, _ := newServer(t, 200, `{"content":"From env","author":"Env"}`)

	withEnv(t, server.URL)

	quote, err := GetQuotable()

	if err != nil || quote.Content != "From env" {
		t.Fatalf("got %+v, %v", quote, err)
	}
}

func TestGetQuotableFromReportsBadStatusAndBadBody(t *testing.T) {
	failing, _ := newServer(t, 503, `{}`)

	if _, err := GetQuotableFrom(failing.URL); err == nil {
		t.Fatal("a 503 must be an error")
	}

	garbage, _ := newServer(t, 200, `not json`)

	if _, err := GetQuotableFrom(garbage.URL); err == nil {
		t.Fatal("a body that is not JSON must be an error")
	}
}
