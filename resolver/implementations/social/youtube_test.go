package social

import (
	"net/http"
	"testing"

	"moe.best.annai/session"
)

var emptyHeaders = http.Header{}

func TestYouTubeGetUrl_SchemeIsHttps(t *testing.T) {
	s := session.NewSession("yt", emptyHeaders)

	url := YOUTUBE().GetUrl(s)

	if url.Scheme != "https" {
		t.Errorf("Scheme = '%s', want 'https'", url.Scheme)
	}
}

func TestYouTubeGetUrl_DefaultLanguageIsEnglish(t *testing.T) {
	s := session.NewSession("yt", emptyHeaders)

	url := YOUTUBE().GetUrl(s)

	if url.Host != "youtube.com" {
		t.Errorf("Host = '%s', want 'youtube.com'", url.Host)
	}
}

func TestYouTubeGetUrl_LanguageProvided_IsIgnored(t *testing.T) {
	headers := http.Header{}
	headers.Add("Accept-Language", "en")
	s := session.NewSession("yt", headers)

	url := YOUTUBE().GetUrl(s)

	if url.Host != "youtube.com" {
		t.Errorf("Host = '%s', want 'youtube.com'", url.Host)
	}
}

func TestYouTubeGetUrl_NoArguments(t *testing.T) {
	s := session.NewSession("yt", emptyHeaders)

	url := YOUTUBE().GetUrl(s)

	if url.Query().Has("query") {
		t.Errorf("Query 'q' = '%s', want nothing", url.Query().Get("q"))
	}
}

func TestYouTubeGetUrl_WithArgument_PathSetToResults(t *testing.T) {
	s := session.NewSession("yt hello", emptyHeaders)

	url := YOUTUBE().GetUrl(s)

	if url.Path != "results" {
		t.Errorf("Path = '%s', want 'results'", url.Path)
	}
}

func TestYouTubeGetUrl_SingleArgument(t *testing.T) {
	s := session.NewSession("yt hello", emptyHeaders)

	url := YOUTUBE().GetUrl(s)

	if !url.Query().Has("search_query") {
		t.Fatalf("Query 'search_query' is not set")
	}
	query := url.Query().Get("search_query")
	if query != "hello" {
		t.Errorf("Query 'search_query' = '%s', want 'hello'", query)
	}
}

func TestYouTubeGetUrl_MultipleArguments(t *testing.T) {
	s := session.NewSession("yt hello world", emptyHeaders)

	url := YOUTUBE().GetUrl(s)

	if !url.Query().Has("search_query") {
		t.Fatalf("Query 'search_query' is not set")
	}
	query := url.Query().Get("search_query")
	if query != "hello world" {
		t.Errorf("Query 'search_query' = '%s', want 'hello world'", query)
	}
}
