package dictionaries

import (
	"net/url"
	"strings"

	"moe.best.annai/request"
	"moe.best.annai/resolver/model"
)

const searchPath = "search"

var baseURL = url.URL{
	Scheme: "https",
	Host:   "jisho.org",
}

func getJishoURL(r request.Request) *url.URL {
	if len(r.Arguments) == 0 {
		return &baseURL
	}

	newUrl := baseURL
	// Append the base path, which is "search"
	newUrl.Path = searchPath
	// Append the user's search query in the path
	newUrl = *newUrl.JoinPath(strings.Join(r.Arguments, " "))

	return &newUrl
}

func JISHO() model.Resolver {
	return model.Resolver{
		GetUrl: getJishoURL,
	}
}
