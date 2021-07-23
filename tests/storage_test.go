package tests

import (
	"log"
	"os"
	"testing"

	tests "github.com/beyondstorage/go-integration-test/v4"
)

func TestStorage(t *testing.T) {
	log.SetOutput(os.Stderr)
	tests.TestStorager(t, setupTest(t))
}
