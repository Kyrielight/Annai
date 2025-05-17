package transit

import (
	"net/url"
	"strings"

	"moe.best.annai/resolver/model"
	"moe.best.annai/session"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "www.google.com",
	Path:   "maps",
}

func getUrl(s session.Session) *url.URL {
	newUrl := baseURL

	if len(s.Arguments) == 0 {
		return &newUrl
	}

	// The search path *ends* with a slash.
	// This is valid: https://www.rfc-editor.org/rfc/rfc3986#section-3.3.
	newUrl = *newUrl.JoinPath("search/")

	search := newUrl.Query()
	search.Set("api", "1")
	search.Set("query", strings.Join(s.Arguments, " "))
	newUrl.RawQuery = search.Encode()

	return &newUrl
}

func GOOGLE_MAPS() model.Resolver {
	return model.Resolver{
		GetUrl: getUrl,
	}
}
