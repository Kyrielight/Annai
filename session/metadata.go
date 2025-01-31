package session

import "net/http"

const acceptLanguage = "Accept-Language"

type Metadata struct {
	Language Language
}

func newMetadata(header http.Header) Metadata {
	return Metadata{
		Language: NewLanguage(header.Get(acceptLanguage)),
	}
}
