package social

import (
	"net/url"
	"strings"

	"golang.org/x/text/language"
	"moe.best.annai/request"
	"moe.best.annai/resolver/model"
)

const searchQueryKey = "search_query"

var baseURL = url.URL{
	Scheme: "https",
	Path:   "results",
}

func createYouTubeURL(host string) url.URL {
	url := baseURL
	url.Host = host
	return url
}

var baseURLs = map[language.Tag]url.URL{
	language.English: createYouTubeURL("youtube.com"),
}

func getYouTubeUrl(r request.Request) *url.URL {
	// Implicit default for this resolver is (no-region English)
	url := baseURLs[language.English]

	if len(r.Arguments) == 0 {
		return &url
	}

	search := url.Query()
	search.Set(searchQueryKey, strings.Join(r.Arguments, " "))
	url.RawQuery = search.Encode()

	return &url
}

func YOUTUBE() model.Resolver {
	return model.Resolver{
		GetUrl: getYouTubeUrl,
	}
}
