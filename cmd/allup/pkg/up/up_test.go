package up

import (
	"log"
	"testing"

	"gopkg.in/src-d/go-git.v4"
)

func TestRepoOpen(t *testing.T) {
	_, err := git.PlainOpen(".")
	log.Printf("open repo, err: %v", err)
}
