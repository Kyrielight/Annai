package dictionaries

import (
	"net/url"
	"strings"

	"moe.best.annai/resolver/model"
	"moe.best.annai/session"
)

const searchPath = "search"

var baseURL = url.URL{
	Scheme: "https",
	Host:   "jisho.org",
}

func getJishoURL(s session.Session) *url.URL {
	if len(s.Arguments) == 0 {
		return &baseURL
	}

	newUrl := baseURL
	// Append the base path, which is "search"
	newUrl.Path = searchPath
	// Append the user's search query in the path
	newUrl = *newUrl.JoinPath(strings.Join(s.Arguments, " "))

	return &newUrl
}

func JISHO() model.Resolver {
	return model.Resolver{
		GetUrl: getJishoURL,
	}
}
