package core

import (
	"net/url"

	"golang.org/x/text/language"

	"moe.best.annai/request"
	"moe.best.annai/resolver/model"
)

var baseURLs = map[language.Tag]url.URL{
	language.BritishEnglish: {Scheme: "https", Host: "google.co.uk"},
	language.English:        {Scheme: "https", Host: "google.com"},
	language.Japanese:       {Scheme: "https", Host: "google.co.jp"},
}

func getUrl(r request.Request) *url.URL {
	// Implicit default for this resolver is (no-region) English.
	url := baseURLs[language.English]
	switch r.Tag {
	case language.Japanese:
		url = baseURLs[language.Japanese]
	}
	return &url
}

func GOOGLE() model.Resolver {
	return model.Resolver{
		GetUrl: getUrl,
	}
}
