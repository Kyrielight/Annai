package core

import (
	"net/http"
	"testing"

	"moe.best.annai/session"
)

var headers = http.Header{}

func TestGoogleGetUrl_SchemeIsHttps(t *testing.T) {
	s := session.NewSession("g", headers)

	url := GOOGLE().GetUrl(s)

	if url.Scheme != "https" {
		t.Errorf("Scheme = '%s', want 'https'", url.Scheme)
	}
}

func TestGoogleGetUrl_LanguageNotSpecified_DefaultsToEnglish(t *testing.T) {
	headers := http.Header{}
	s := session.NewSession("g", headers)

	url := GOOGLE().GetUrl(s)

	if url.Host != "google.com" {
		t.Errorf("Host = '%s', want 'google.com'", url.Host)
	}
}

func TestGoogleGetUrl_English(t *testing.T) {
	headers := http.Header{}
	headers.Add("Accept-Language", "en")
	s := session.NewSession("g", headers)

	url := GOOGLE().GetUrl(s)

	if url.Host != "google.com" {
		t.Errorf("Host = '%s', want 'google.com'", url.Host)
	}
}

func TestGoogleGetUrl_BritishEnglish(t *testing.T) {
	headers := http.Header{}
	headers.Add("Accept-Language", "en-uk")
	s := session.NewSession("g", headers)

	url := GOOGLE().GetUrl(s)

	if url.Host != "google.co.uk" {
		t.Errorf("Host = '%s', want 'google.co.uk'", url.Host)
	}
}

func TestGoogleGetUrl_Japanese(t *testing.T) {
	headers := http.Header{}
	headers.Add("Accept-Language", "ja")
	s := session.NewSession("g", headers)

	url := GOOGLE().GetUrl(s)

	if url.Host != "google.co.jp" {
		t.Errorf("Host = '%s', want 'google.co.jp'", url.Host)
	}
}

func TestGoogleGetUrl_NoArguments(t *testing.T) {
	request := session.NewSession("g", headers)

	url := GOOGLE().GetUrl(request)

	if url.Query().Has("query") {
		t.Errorf("Query 'q' = '%s', want nothing", url.Query().Get("q"))
	}
}

func TestGoogleGetUrl_WithArgument_PathSetToSearch(t *testing.T) {
	request := session.NewSession("g hello", headers)

	url := GOOGLE().GetUrl(request)

	if url.Path != "search" {
		t.Errorf("Path = '%s', want 'search'", url.Path)
	}
}

func TestGoogleGetUrl_SingleArgument(t *testing.T) {
	request := session.NewSession("g hello", headers)

	url := GOOGLE().GetUrl(request)

	if !url.Query().Has("q") {
		t.Fatalf("Query 'q' is not set")
	}
	query := url.Query().Get("q")
	if query != "hello" {
		t.Errorf("Query 'q' = '%s', want 'hello'", query)
	}
}
func TestGoogleGetUrl_MultipleArguments(t *testing.T) {
	request := session.NewSession("g hello world", headers)

	url := GOOGLE().GetUrl(request)

	if !url.Query().Has("q") {
		t.Fatalf("Query 'q' is not set")
	}
	query := url.Query().Get("q")
	if query != "hello world" {
		t.Errorf("Query 'q' = '%s', want 'hello world'", query)
	}
}
