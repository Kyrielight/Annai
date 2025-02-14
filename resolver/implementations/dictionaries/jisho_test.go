package dictionaries

import (
	"net/http"
	"strings"
	"testing"

	"moe.best.annai/session"
)

var emptyHeaders = http.Header{}

func TestJishoGetUrl_Scheme(t *testing.T) {
	session := session.NewSession("jisho", emptyHeaders)

	url := JISHO().GetUrl(session)

	if url.Scheme != "https" {
		t.Errorf("Scheme = '%s', want 'https'", url.Scheme)
	}
}

func TestJishoGetUrl_Host(t *testing.T) {
	session := session.NewSession("jisho", emptyHeaders)

	url := JISHO().GetUrl(session)

	if url.Host != "jisho.org" {
		t.Errorf("Host = '%s', wanted 'jisho.org'", url.Path)
	}
}

func TestJishoGetUrl_LanguageDisregarded_UsesEnglish(t *testing.T) {
	headers := http.Header{}
	headers.Add("Accept-Language", "ja")
	session := session.NewSession("jisho hello world", headers)

	url := JISHO().GetUrl(session)

	if url.Host != "jisho.org" {
		t.Errorf("Host = '%s', wanted 'jisho.org'", url.Path)
	}
}

func TestJishoGetUrl_NoArguments(t *testing.T) {
	session := session.NewSession("jisho", emptyHeaders)

	url := JISHO().GetUrl(session)

	if len(url.Path) > 0 {
		t.Errorf("Path = '%s', wanted nothing", url.Path)
	}
}

func TestJishoGetUrl_WithArgument_PathStartsWithSearch(t *testing.T) {
	session := session.NewSession("jisho hello world", emptyHeaders)

	url := JISHO().GetUrl(session)

	path := strings.Split(url.Path, "/")[0]
	if path != "search" {
		t.Errorf("First path is '%s', wanted 'search'", url.Path)
	}
}

func TestJishoGetUrl_SingleArgument(t *testing.T) {
	session := session.NewSession("jisho hello", emptyHeaders)

	url := JISHO().GetUrl(session)

	argument := strings.Split(url.Path, "/")[1]
	if argument != "hello" {
		t.Errorf("Argument is '%s', wanted 'hello'", argument)
	}
}

func TestJishoGetUrl_MultipleArguments(t *testing.T) {
	session := session.NewSession("jisho hello world", emptyHeaders)

	url := JISHO().GetUrl(session)

	arguments := strings.Split(url.Path, "/")[1]
	if arguments != "hello world" {
		t.Errorf("Arguments are '%s', wanted 'hello world'", arguments)
	}
}
