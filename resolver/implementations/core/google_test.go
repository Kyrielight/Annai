package core

import (
	"testing"

	"moe.best.annai/request"
)

func TestGetUrl_SchemeIsHttps(t *testing.T) {
	request := request.NewRequest("g")

	url := GOOGLE().GetUrl(request)

	if url.Scheme != "https" {
		t.Errorf("Scheme = %s, want 'https'", url.Scheme)
	}
}

func TestGetUrl_DefaultLanguageIsEnglish(t *testing.T) {
	request := request.NewRequest("g")

	url := GOOGLE().GetUrl(request)

	if url.Host != "google.com" {
		t.Errorf("Host = %s, want 'google.com'", url.Host)
	}
}

func TestGetUrl_NoArguments(t *testing.T) {
	request := request.NewRequest("g")

	url := GOOGLE().GetUrl(request)

	if url.Query().Has("query") {
		t.Errorf("Query 'q' = %s, want nothing", url.Query().Get("q"))
	}
}

func TestGetUrl_WithArgument_PathSetToSearch(t *testing.T) {
	request := request.NewRequest("g hello")

	url := GOOGLE().GetUrl(request)

	if url.Path != "search" {
		t.Errorf("Path = %s, want 'search'", url.Path)
	}
}

func TestGetUrl_SingleArgument(t *testing.T) {
	request := request.NewRequest("g hello")

	url := GOOGLE().GetUrl(request)

	if !url.Query().Has("q") {
		t.Fatalf("Query 'q' is not set")
	}
	query := url.Query().Get("q")
	if query != "hello" {
		t.Errorf("Query 'q' = %s, want 'hello", query)
	}
}
func TestGetUrl_MultipleArguments(t *testing.T) {
	request := request.NewRequest("g hello world")

	url := GOOGLE().GetUrl(request)

	if !url.Query().Has("q") {
		t.Fatalf("Query 'q' is not set")
	}
	query := url.Query().Get("q")
	if query != "hello world" {
		t.Errorf("Query 'q' = %s, want 'hello world'", query)
	}
}
