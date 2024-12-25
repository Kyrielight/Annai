package core

import (
	"net/url"
	"strings"

	"golang.org/x/text/language"

	"moe.best.annai/request"
	"moe.best.annai/resolver/model"
)

var baseURLs = map[language.Tag]url.URL{
	language.BritishEnglish: {Scheme: "https", Host: "google.co.uk"},
	language.English:        {Scheme: "https", Host: "google.com"},
	language.Japanese:       {Scheme: "https", Host: "google.co.jp"},
}

func getGoogleUrl(r request.Request) *url.URL {
	// Implicit default for this resolver is (no-region) English.
	url := baseURLs[language.English]

	if localUrl, exists := baseURLs[r.Tag]; exists {
		url = localUrl
	}

	if len(r.Arguments) == 0 {
		return &url
	}

	url.Path = "search"
	search := url.Query()
	search.Set("q", strings.Join(r.Arguments, " "))
	url.RawQuery = search.Encode()

	return &url
}

func GOOGLE() model.Resolver {
	return model.Resolver{
		GetUrl: getGoogleUrl,
	}
}
