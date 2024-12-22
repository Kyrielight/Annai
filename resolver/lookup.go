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
	return defaultResolver.GetUrl(r)
}
