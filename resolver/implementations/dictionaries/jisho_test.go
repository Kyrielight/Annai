package dictionaries

import (
	"strings"
	"testing"

	"moe.best.annai/request"
)

func TestJishoGetUrl_Scheme(t *testing.T) {
	request := request.NewRequest("jisho")

	url := JISHO().GetUrl(request)

	if url.Scheme != "https" {
		t.Errorf("Scheme = '%s', want 'https'", url.Scheme)
	}
}

func TestJishoGetUrl_Host(t *testing.T) {
	request := request.NewRequest("jisho")

	url := JISHO().GetUrl(request)

	if url.Host != "jisho.org" {
		t.Errorf("Host = '%s', wanted 'jisho.org'", url.Path)
	}
}

func TestJishoGetUrl_NoArguments(t *testing.T) {
	request := request.NewRequest("jisho")

	url := JISHO().GetUrl(request)

	if len(url.Path) > 0 {
		t.Errorf("Path = '%s', wanted nothing", url.Path)
	}
}

func TestJishoGetUrl_WithArgument_PathStartsWithSearch(t *testing.T) {
	request := request.NewRequest("jisho hello world")

	url := JISHO().GetUrl(request)

	path := strings.Split(url.Path, "/")[0]
	if path != "search" {
		t.Errorf("First path is '%s', wanted 'search'", url.Path)
	}
}

func TestJishoGetUrl_SingleArgument(t *testing.T) {
	request := request.NewRequest("jisho hello")

	url := JISHO().GetUrl(request)

	argument := strings.Split(url.Path, "/")[1]
	if argument != "hello" {
		t.Errorf("Argument is '%s', wanted 'hello'", argument)
	}
}

func TestJishoGetUrl_MultipleArguments(t *testing.T) {
	request := request.NewRequest(("jisho hello world"))

	url := JISHO().GetUrl(request)

	arguments := strings.Split(url.Path, "/")[1]
	if arguments != "hello world" {
		t.Errorf("Arguments are '%s', wanted 'hello world'", arguments)
	}
}
