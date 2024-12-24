package resolver

import (
	"net/url"

	"moe.best.annai/request"
	"moe.best.annai/resolver/implementations/core"
)

var defaultResolver = core.GOOGLE()

func Lookup(r request.Request) *url.URL {
	if resolver, ok := resolvers[r.Command]; ok {
		return resolver.GetUrl(r)
	}
	// TODO: Process defaults differently - the command should be included
	// in arguments as the Google resolver is implicit.
	return defaultResolver.GetUrl(r)
}
