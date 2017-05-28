package monobullet

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNote(t *testing.T) {
	ConfigFromFile()
	config.Debug = true
	push := new(Push)
	push.Type = NoteType
	push.Title = "test title"
	push.Body = "test body"

	resp, err := sendPush(push)
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

func TestLive(t *testing.T) {
	ConfigFromFile()
	config.Debug = true
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		wsConnect(ctx)
	}(ctx)

	push := new(Push)
	push.Type = NoteType
	push.Title = "test title"
	push.Body = "test body"

	go func() {
		var err error
		push, err = sendPush(push)
		if err != nil {
			t.Error(err)
		}
	}()

Poll:
	for {
		select {
		case note := <-PushChannel:
			if push.Iden == note.Iden {
				fmt.Printf("got push with correct iden %v\n", push.Iden)
				break Poll
			}
		case <-ctx.Done():
			t.Error("context cancelled before realtime event received")
		}
	}
}
