package social

import (
	"testing"

	"moe.best.annai/request"
)

func TestYouTubeGetUrl_Scheme(t *testing.T) {
	request := request.NewRequest("yt")

	url := YOUTUBE().GetUrl(request)

	if url.Scheme != "https" {
		t.Errorf("Scheme = '%s', want 'https'", url.Scheme)
	}
}
func TestYouTubeGetUrl_DefaultLanguageIsEnglish(t *testing.T) {
	request := request.NewRequest("yt")

	url := YOUTUBE().GetUrl(request)

	if url.Host != "youtube.com" {
		t.Errorf("Host = '%s', want 'youtube.com'", url.Host)
	}
}

func TestYouTubeGetUrl_NoArguments(t *testing.T) {
	request := request.NewRequest("yt")

	url := YOUTUBE().GetUrl(request)

	if url.Query().Has("query") {
		t.Errorf("Query 'q' = '%s', want nothing", url.Query().Get("q"))
	}
}

func TestYouTubeGetUrl_WithArgument_PathSetToResults(t *testing.T) {
	request := request.NewRequest("yt hello")

	url := YOUTUBE().GetUrl(request)

	if url.Path != "results" {
		t.Errorf("Path = '%s', want 'results'", url.Path)
	}
}

func TestYouTubeGetUrl_SingleArgument(t *testing.T) {
	request := request.NewRequest("yt hello")

	url := YOUTUBE().GetUrl(request)

	if !url.Query().Has("search_query") {
		t.Fatalf("Query 'search_query' is not set")
	}
	query := url.Query().Get("search_query")
	if query != "hello" {
		t.Errorf("Query 'search_query' = '%s', want 'hello'", query)
	}
}

func TestYouTubeGetUrl_MultipleArguments(t *testing.T) {
	request := request.NewRequest("yt hello world")

	url := YOUTUBE().GetUrl(request)

	if !url.Query().Has("search_query") {
		t.Fatalf("Query 'search_query' is not set")
	}
	query := url.Query().Get("search_query")
	if query != "hello world" {
		t.Errorf("Query 'search_query' = '%s', want 'hello world'", query)
	}
}
