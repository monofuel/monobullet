package monobullet

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	ConfigFromFile()
}

func TestSelf(t *testing.T) {
	ConfigFromFile()
	config.Debug = true

	user, err := getUser()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Fetched user: %v\n", user)
}
