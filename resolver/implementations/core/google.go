package core

import (
	"net/url"
	"strings"

	"golang.org/x/text/language"

	"moe.best.annai/resolver/model"
	"moe.best.annai/session"
)

var baseURLs = map[language.Tag]url.URL{
	language.BritishEnglish: {Scheme: "https", Host: "google.co.uk"},
	language.English:        {Scheme: "https", Host: "google.com"},
	language.Japanese:       {Scheme: "https", Host: "google.co.jp"},
}

func getGoogleUrl(s session.Session) *url.URL {

	url, _ := s.Metadata.Language.MatchUrlMap(baseURLs, baseURLs[language.English])

	if len(s.Arguments) == 0 {
		return &url
	}

	url.Path = "search"
	search := url.Query()
	search.Set("q", strings.Join(s.Arguments, " "))
	url.RawQuery = search.Encode()

	return &url
}

func GOOGLE() model.Resolver {
	return model.Resolver{
		GetUrl: getGoogleUrl,
	}
}
