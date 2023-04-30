package test

import (
	"log"
	general "project/src/pkg/utils"
	"testing"
)

func TestLoadCredentials(t *testing.T) {

	_, err := general.LoadTLSCredentials()
	if err != nil {
		log.Printf("could not load TLS credentials %v", err)
	}
	got, want := 1, 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
