package social

import (
	"net/url"
	"strings"

	"golang.org/x/text/language"
	"moe.best.annai/resolver/model"
	"moe.best.annai/session"
)

const searchQueryKey = "search_query"

func createYouTubeURL(host string) url.URL {
	url := url.URL{
		Scheme: "https",
		Path:   "results",
	}
	url.Host = host
	return url
}

var baseURLs = map[language.Tag]url.URL{
	language.English: createYouTubeURL("youtube.com"),
}

func getYouTubeUrl(s session.Session) *url.URL {
	// Implicit default for this resolver is (no-region English)
	url := baseURLs[language.English]

	if len(s.Arguments) == 0 {
		return &url
	}

	search := url.Query()
	search.Set(searchQueryKey, strings.Join(s.Arguments, " "))
	url.RawQuery = search.Encode()

	return &url
}

func YOUTUBE() model.Resolver {
	return model.Resolver{
		GetUrl: getYouTubeUrl,
	}
}
