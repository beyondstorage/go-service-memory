package tests

import (
	"testing"

	tests "github.com/beyondstorage/go-integration-test/v4"
)

func TestStorage(t *testing.T) {
	tests.TestStorager(t, setupTest(t))
}

func TestAppend(t *testing.T) {
	tests.TestAppender(t, setupTest(t))
}

func TestDir(t *testing.T) {
	tests.TestDirer(t, setupTest(t))
}

func TestCopy(t *testing.T) {
	tests.TestCopier(t, setupTest(t))
}

func TestMove(t *testing.T) {
	tests.TestMover(t, setupTest(t))
}
