package session

import (
	"net/http"
	"strings"
)

type Session struct {
	Command   string
	Arguments []string
	Metadata  Metadata
}

func NewSession(query string, header http.Header) Session {
	if len(query) == 0 {
		return Session{}
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

	return Session{
		Command:   command,
		Arguments: arguments,
		Metadata:  newMetadata(header),
	}
}
