package resolver

import (
	"moe.best.annai/resolver/implementations/core"
	"moe.best.annai/resolver/implementations/dictionaries"
	"moe.best.annai/resolver/implementations/social"
	"moe.best.annai/resolver/implementations/transit"
	"moe.best.annai/resolver/model"
)

var resolvers = map[string]model.Resolver{
	"g":       core.GOOGLE(),
	"gmaps":   transit.GOOGLE_MAPS(),
	"google":  core.GOOGLE(),
	"j":       dictionaries.JISHO(),
	"jisho":   dictionaries.JISHO(),
	"maps":    transit.GOOGLE_MAPS(),
	"youtube": social.YOUTUBE(),
	"yt":      social.YOUTUBE(),
}
