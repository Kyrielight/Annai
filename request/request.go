package request

import "golang.org/x/text/language"

type Request struct {
	Command   string
	Arguments []string
	Tag       language.Tag
}

func NewRequest(command string, arguments []string, tag language.Tag) Request {
	return Request{}
}

func NewRequestWithoutArguments(command string, tag language.Tag) Request {
	return Request{
		Command:   command,
		Arguments: []string{},
		Tag:       tag,
	}
}
