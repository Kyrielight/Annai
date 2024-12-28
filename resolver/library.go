package resolver

import (
	"moe.best.annai/resolver/implementations/core"
	"moe.best.annai/resolver/implementations/dictionaries"
	"moe.best.annai/resolver/implementations/social"
	"moe.best.annai/resolver/model"
)

var resolvers = map[string]model.Resolver{
	"g":       core.GOOGLE(),
	"google":  core.GOOGLE(),
	"j":       dictionaries.JISHO(),
	"jisho":   dictionaries.JISHO(),
	"youtube": social.YOUTUBE(),
	"yt":      social.YOUTUBE(),
}
