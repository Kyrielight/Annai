package model

import (
	"net/url"

	"moe.best.annai/request"
)

type Resolver struct {
	GetUrl func(request.Request) *url.URL
}
