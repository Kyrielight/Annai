package request

import (
	"strings"
	"testing"
)

func TestNewRequest_TargetCommandOnly(t *testing.T) {
	query := "google"

	request := NewRequest(query)

	if (request.Command) != query {
		t.Errorf("Command = %s, want '%s'", request.Command, query)
	}
	if len(request.Arguments) != 0 {
		t.Errorf("Arguments = %s, want nothing", strings.Join(request.Arguments, " "))
	}
}

func TestNewRequest_SlashRequestOnly(t *testing.T) {
	query := "google/"

	request := NewRequest(query)

	if (request.Command) != "google" {
		t.Errorf("Command = %s, want '%s'", request.Command, "google")
	}
	if len(request.Arguments) != 0 {
		t.Errorf("Arguments = %s, want nothing", strings.Join(request.Arguments, " "))
	}
}

func TestNewRequest_TargetCommandWithSingleArgument(t *testing.T) {
	query := "google hello"

	request := NewRequest(query)

	if (request.Command) != "google" {
		t.Errorf("Command = %s, want 'google'", request.Command)
	}
	if len(request.Arguments) != 1 {
		t.Fatalf("# of Arguments = %d, wanted 1", len(request.Arguments))
	}
	if request.Arguments[0] != "hello" {
		t.Errorf("Arguments = %s, want 'hello'", strings.Join(request.Arguments, " "))
	}
}

func TestNewRequest_SlashCommandWithSingleArgument(t *testing.T) {
	query := "google/hello"

	request := NewRequest(query)

	if (request.Command) != "google" {
		t.Errorf("Command = %s, want 'google'", request.Command)
	}
	if len(request.Arguments) != 1 {
		t.Fatalf("# of Arguments = %d, wanted 1", len(request.Arguments))
	}
	if request.Arguments[0] != "hello" {
		t.Errorf("Arguments = %s, want 'hello'", strings.Join(request.Arguments, " "))
	}
}

func TestNewRequest_TargetCommandWithMultipleArguments(t *testing.T) {
	query := "google hello world"

	request := NewRequest(query)

	if (request.Command) != "google" {
		t.Errorf("Command = %s, want 'google'", request.Command)
	}
	if len(request.Arguments) != 2 {
		t.Fatalf("# of Arguments = %d, wanted 2", len(request.Arguments))
	}
	if request.Arguments[0] != "hello" {
		t.Errorf("Arguments = %s, want 'hello'", strings.Join(request.Arguments, " "))
	}
	if request.Arguments[1] != "world" {
		t.Errorf("Arguments = %s, want 'world'", strings.Join(request.Arguments, " "))
	}
}

func TestNewRequest_SlashCommandWithMultipleArguments(t *testing.T) {
	query := "google/hello world"

	request := NewRequest(query)

	if (request.Command) != "google" {
		t.Errorf("Command = %s, want 'google'", request.Command)
	}
	if len(request.Arguments) != 2 {
		t.Fatalf("# of Arguments = %d, wanted 2", len(request.Arguments))
	}
	if request.Arguments[0] != "hello" {
		t.Errorf("Arguments = %s, want 'hello'", strings.Join(request.Arguments, " "))
	}
	if request.Arguments[1] != "world" {
		t.Errorf("Arguments = %s, want 'world'", strings.Join(request.Arguments, " "))
	}
}
