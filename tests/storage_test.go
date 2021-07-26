package tests

import (
	"testing"

	tests "github.com/beyondstorage/go-integration-test/v4"
)

func TestStorage(t *testing.T) {
	tests.TestStorager(t, setupTest(t))
}
