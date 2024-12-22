package resolver

import (
	"moe.best.annai/resolver/implementations/core"
	"moe.best.annai/resolver/model"
)

var resolvers = map[string]model.Resolver{
	"g":      core.GOOGLE(),
	"google": core.GOOGLE(),
}
