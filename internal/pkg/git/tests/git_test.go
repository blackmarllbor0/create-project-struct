package tests

import (
	"fmt"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/git"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/git/tests/mock"
	"os"
	"testing"
)

//TODO the tests run and run fine individually, but if you run them together,
// you get confused with the test directories.

func TestCreateLocalGitRepository_CurrentDir(t *testing.T) {
	t.Log("create dir for test")
	{
		if err := os.Mkdir(mock.TestDir, 0755); err != nil {
			t.Fatalf("Failed to create test directory: %v", err)
		}
		defer os.RemoveAll(fmt.Sprintf("../%s", mock.TestDir))

		if err := os.Chdir(mock.TestDir); err != nil {
			t.Fatalf("Expected no error, but got: %s", err)
		}

		defer func() {
			if err := os.Chdir(".."); err != nil {
				t.Fatalf("Expected no error, but got: %s", err)
			}
		}()
	}

	t.Log("create repository in current dir")
	{
		if err := git.CreateLocalGitRepository(mock.IsCurrentDirIsTrue{}); err != nil {
			t.Fatalf("Expected no error, but got: %s", err)
		}
	}

	t.Log("check the .git in the current directory")
	{
		_, err := os.Stat(".git")
		if os.IsNotExist(err) {
			t.Fatal(".git directory not found in the expected location")
		}
	}

	t.Log("SUCCESS")
}

func TestCreateLocalGitRepository_NewDir(t *testing.T) {
	t.Log("create dir for test")
	{
		if err := os.Mkdir(mock.TestDir, 0755); err != nil {
			t.Fatalf("Failed to create test directory: %v", err)
		}
		defer os.RemoveAll(mock.TestDir)
	}

	t.Log("create repository in new dir")
	{
		if err := git.CreateLocalGitRepository(mock.IsCurrentDirIsTrue{}); err != nil {
			t.Fatalf("Expected no error, but got: %s", err)
		}
	}

	t.Log("check the .git in the current directory")
	{
		_, err := os.Stat(".git")
		if os.IsNotExist(err) {
			t.Fatal(".git directory not found in the expected location")
		}
	}

	t.Log("SUCCESS")
}
