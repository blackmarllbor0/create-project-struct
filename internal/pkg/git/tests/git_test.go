package tests

import (
	"github.com/blackmarllboro/create-project-struct/internal/pkg/git"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/git/tests/mock"

	"os"
	"testing"
)

// TODO we need to figure out how to create and navigate to the test dirs.
func TestCreateLocalGitRepository_CurrentDir(t *testing.T) {
	if err := os.Mkdir(mock.TestDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(mock.TestDir)

	if err := os.Chdir(mock.TestDir); err != nil {
		t.Fatalf("Expected no error, but got: %s", err)
	}

	if err := git.CreateLocalGitRepository(mock.IsCurrentDirIsTrue{}); err != nil {
		t.Fatalf("Expected no error, but got: %s", err)
	}

	_, err := os.Stat(mock.TestDir)
	if os.IsNotExist(err) {
		t.Fatal(".git directory not found in the expected location")
	}
}
