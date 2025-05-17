package transit

import (
	"net/http"
	"testing"

	"moe.best.annai/session"
)

var emptyHeaders = http.Header{}

func TestGoogleMapsGetUrl_SchemeIsHttps(t *testing.T) {
	s := session.NewSession("gmaps", emptyHeaders)

	url := GOOGLE_MAPS().GetUrl(s)

	if url.Scheme != "https" {
		t.Errorf("Scheme = '%s', want 'https'", url.Scheme)
	}
}

func TestGoogleMapsGetUrl_LanguageProvided_IsIgnored(t *testing.T) {
	headers := http.Header{}
	headers.Add("Accept-Language", "en")
	s := session.NewSession("gmaps", headers)

	url := GOOGLE_MAPS().GetUrl(s)

	if url.Host != "www.google.com" {
		t.Errorf("Host = '%s', want 'www.google.com'", url.Host)
	}
}

func TestGoogleMapsGetUrl_NoArguments(t *testing.T) {
	const expectedUrl = "https://www.google.com/maps"
	s := session.NewSession("gmaps", emptyHeaders)

	url := GOOGLE_MAPS().GetUrl(s)

	if url.String() != expectedUrl {
		t.Errorf("URL String = '%s', want '%s'", url.String(), expectedUrl)
	}
}

func TestGoogleMapsGetUrl_WithArgument_QueryHasApiSet(t *testing.T) {
	s := session.NewSession("gmaps Shibuya Station", emptyHeaders)

	url := GOOGLE_MAPS().GetUrl(s)

	query := url.Query().Get("api")
	if query != "1" {
		t.Errorf("Query 'api' = '%s', want '1'", query)
	}
}

func TestGoogleMapsGetUrl_SingleArgument(t *testing.T) {
	s := session.NewSession("gmaps Shibuya", emptyHeaders)

	url := GOOGLE_MAPS().GetUrl(s)

	query := url.Query().Get("query")
	if query == "" {
		t.Fatalf("Query 'query' is not set")
	}
	if query != "Shibuya" {
		t.Errorf("Query 'query' = '%s', want 'Shibuya'", query)
	}
}

func TestGoogleMapsGetUrl_MultipleArguments(t *testing.T) {
	s := session.NewSession("gmaps Shibuya Station", emptyHeaders)

	url := GOOGLE_MAPS().GetUrl(s)

	query := url.Query().Get("query")
	if query == "" {
		t.Fatalf("Query 'query' is not set")
	}
	if query != "Shibuya Station" {
		t.Errorf("Query 'query' = '%s', want 'Shibuya Station'", query)
	}
}
