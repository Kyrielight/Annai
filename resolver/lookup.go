package resolver

import (
	"net/url"

	"moe.best.annai/resolver/implementations/core"
	"moe.best.annai/session"
)

var defaultResolver = core.GOOGLE()

func Lookup(s session.Session) *url.URL {

	if resolver, ok := resolvers[s.Command]; ok {
		return resolver.GetUrl(s)
	}
	// TODO: Process defaults differently - the command should be included
	// in arguments as the Google resolver is implicit.
	defaultRequest := session.Session{
		Command:   "default",
		Arguments: append([]string{s.Command}, s.Arguments...),
	}

	return defaultResolver.GetUrl(defaultRequest)
}
