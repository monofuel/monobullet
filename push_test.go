package monobullet

import (
	"fmt"
	"testing"
)

func TestNote(t *testing.T) {
	ConfigFromFile()
	config.Debug = true
	push := new(Note)
	push.Type = NoteType
	push.Title = "test title"
	push.Body = "test body"

	resp, err := sendNote(push)
	if err != nil {
		t.Error(err)
	}
	if resp.Type != push.Type {
		t.Errorf("invalid type, expected %v got %v", push.Type, resp.Type)
	}
	if resp.Title != push.Title {
		t.Errorf("invalid title, expected %v got %v", push.Title, resp.Title)
	}
	if resp.Body != push.Body {
		t.Errorf("invalid body, expected %v got %v", push.Body, resp.Body)
	}
}

func TestSelf(t *testing.T) {
	ConfigFromFile()
	config.Debug = true

	user, err := getUser()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Fetched user: %v\n", user)
	if user.Iden == "" {
		t.Error("missing iden")
	}
	if user.Email == "" {
		t.Error("missing email")
	}
	if user.Created == 0 {
		t.Error("missing when created")
	}
}
