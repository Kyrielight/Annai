package session

import (
	"net/url"
	"testing"

	"golang.org/x/text/language"
)

var _ENGLISH_URL = url.URL{Host: "google.com"}
var _JAPANESE_URL = url.URL{Host: "google.co.jp"}

func TestLanguage_MatchTag_AcceptNotInSupportedTags_DefaultInSupportedTags_UsesDefault(t *testing.T) {
	supportedTags := []language.Tag{language.English, language.Japanese}
	l := NewLanguage(language.Arabic.String())

	match := l.MatchTag(supportedTags, language.English)

	if match != language.English {
		t.Fatalf("Matched to '%s', wanted '%s'", match.String(), language.English.String())
	}

}

func TestLanguage_MatchTag_AcceptNotInSupportedTags_DefaultNotInSupportedTags_UsesDefault(t *testing.T) {
	supportedTags := []language.Tag{language.English, language.Japanese}
	l := NewLanguage(language.Arabic.String())

	match := l.MatchTag(supportedTags, language.Chinese)

	if match != language.Chinese {
		t.Fatalf("Matched to '%s', wanted '%s'", match.String(), language.Chinese.String())
	}

}

func TestLanguage_MatchTag_AcceptInSupportedTags_UsesAccept(t *testing.T) {
	supportedTags := []language.Tag{language.Japanese}
	l := NewLanguage(language.Japanese.String())

	match := l.MatchTag(supportedTags, language.English)

	if match != language.Japanese {
		t.Fatalf("Matched to '%s', wanted '%s'", match.String(), language.Japanese.String())
	}
}

func TestLanguage_MatchUrlMap_AcceptNotInSupportedTags_UsesDefaultUrl(t *testing.T) {
	supportedTags := map[language.Tag]url.URL{language.English: _ENGLISH_URL}
	l := NewLanguage(language.Japanese.String())

	match, found := l.MatchUrlMap(supportedTags, _JAPANESE_URL)

	if match != _JAPANESE_URL {
		t.Errorf("Matched to '%s', wanted '%s'", match.String(), _JAPANESE_URL.String())
	}
	if found {
		t.Errorf("Unsupported Accept-Language string reported as 'found'")
	}
}

func TestLanguage_MatchUrlMap_AcceptInSupportedTags_UsesAccept(t *testing.T) {
	supportedTags := map[language.Tag]url.URL{
		language.English:  _ENGLISH_URL,
		language.Japanese: _JAPANESE_URL,
	}
	l := NewLanguage(language.Japanese.String())

	match, found := l.MatchUrlMap(supportedTags, _JAPANESE_URL)

	if match != _JAPANESE_URL {
		t.Errorf("Matched to '%s', wanted '%s'", match.String(), _JAPANESE_URL.String())
	}
	if !found {
		t.Errorf("Supported Accept-Language string reported as 'not found'")
	}

}
