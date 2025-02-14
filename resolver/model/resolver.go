package model

import (
	"net/url"

	"moe.best.annai/session"
)

type Resolver struct {
	GetUrl func(session.Session) *url.URL
}
