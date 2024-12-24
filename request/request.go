package request

import (
	"strings"

	"golang.org/x/text/language"
)

type Request struct {
	Command   string
	Arguments []string
	Tag       language.Tag
}

func NewRequest(query string) Request {
	if len(query) == 0 {
		return Request{}
	}

	arguments := []string{}
	rawArguments := strings.Split(query, " ")

	command, firstArgument, found := strings.Cut(rawArguments[0], "/")

	// If the first word ended in a "/", the substring after the slash (if it exists)
	// is considered the first argument.
	if found && len(firstArgument) > 0 {
		arguments = append(arguments, firstArgument)
	}

	// Append any remaining arguments.
	if len(rawArguments) > 1 {
		arguments = append(arguments, rawArguments[1:]...)
	}

	return Request{
		Command:   command,
		Arguments: arguments,
		Tag:       language.English,
	}
}
