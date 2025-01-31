package session

import (
	"net/url"

	gl "golang.org/x/text/language"
)

type Language struct {
	// Provided a list of Tags, return the closest associated Tag. If no Tag is a match, returns the default Tag.
	MatchTag func(supportedTags []gl.Tag, defaultTag gl.Tag) (tag gl.Tag)
	// Provided a map of Tag -> URLs, return the URL associated with the closest Tag. If no Tag is an (exact) match,
	// return the default URL.
	MatchUrlMap func(supportedTags map[gl.Tag]url.URL, defaultUrl url.URL) (redirectURL url.URL, found bool)
}

// Create a Language struct.
// Accepts an Accept-Language header string value.
func NewLanguage(acceptLanguage string) Language {
	matchTag := func(supportedTags []gl.Tag, defaultTag gl.Tag) gl.Tag {

		// First tag is default. (There can be duplicate tags, but it doesn't change the return).
		supportedTagsWithDefault := append([]gl.Tag{defaultTag}, supportedTags...)

		matcher := gl.NewMatcher(supportedTagsWithDefault)
		tag, _ := gl.MatchStrings(matcher, acceptLanguage)

		return tag
	}

	matchUrlMap := func(supportedTags map[gl.Tag]url.URL, defaultURL url.URL) (redirectURL url.URL, found bool) {

		acceptLanguageTags, _, error := gl.ParseAcceptLanguage(acceptLanguage)
		if error != nil {
			return defaultURL, false
		}

		for _, acceptLanguageTag := range acceptLanguageTags {
			if mapURL, exists := supportedTags[acceptLanguageTag]; exists {
				return mapURL, true
			}
		}

		return defaultURL, false
	}

	return Language{
		MatchTag:    matchTag,
		MatchUrlMap: matchUrlMap,
	}
}
