package resolver

import (
	"bytes"
	"net/http"
	"testing"

	"moe.best.annai/resolver/implementations/core"
	"moe.best.annai/resolver/implementations/dictionaries"
	"moe.best.annai/resolver/implementations/social"
	"moe.best.annai/resolver/implementations/transit"
	"moe.best.annai/resolver/model"
	"moe.best.annai/session"
)

func TestLookup_Default_CommandOnly_isGoogle(t *testing.T) {
	testLookup("badcommand", core.GOOGLE(), t)
}

func TestLookup_Default_CommandWithArguments_isGoogle(t *testing.T) {
	testLookup("badcommand jam", core.GOOGLE(), t)
}

func TestLookup_g_isGoogle(t *testing.T) {
	testLookup("g", core.GOOGLE(), t)
}

func TestLookup_gmaps_isGoogleMaps(t *testing.T) {
	testLookup("gmaps", transit.GOOGLE_MAPS(), t)
}

func TestLookup_google_isGoogle(t *testing.T) {
	testLookup("google", core.GOOGLE(), t)
}

func TestLookup_j_isJisho(t *testing.T) {
	testLookup("j", dictionaries.JISHO(), t)
}

func TestLookup_jisho_isJisho(t *testing.T) {
	testLookup("jisho", dictionaries.JISHO(), t)
}

func TestLookup_maps_isGoogleMaps(t *testing.T) {
	testLookup("maps", transit.GOOGLE_MAPS(), t)
}

func TestLookup_youtube_isYouTube(t *testing.T) {
	testLookup("youtube", social.YOUTUBE(), t)
}

func TestLookup_yt_isYouTube(t *testing.T) {
	testLookup("yt", social.YOUTUBE(), t)
}

// Helper to assert a command points to a resolver.
func testLookup(command string, expectedResolver model.Resolver, t *testing.T) {
	s := session.NewSession(command, http.Header{})

	url := Lookup(s)
	expectedUrl := expectedResolver.GetUrl(s)

	urlMarshal, error := url.MarshalBinary()
	if error != nil {
		t.Fatalf("%s", error.Error())
	}
	expectedUrlMarshal, error := url.MarshalBinary()
	if error != nil {
		t.Fatalf("%s", error.Error())
	}
	if !bytes.Equal(urlMarshal, expectedUrlMarshal) {
		t.Fatalf("Url = '%s', wanted '%s'", url.String(), expectedUrl.String())
	}
}
