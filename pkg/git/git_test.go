package git

import (
	"fmt"
	"os"
	"testing"

	"github.com/blackmarllboro/create-project-struct/internal/pkg/args/interfaces"

	"github.com/blackmarllboro/create-project-struct/internal/pkg/args/mocks"
)

func TestCreateLocalGitRepository(t *testing.T) {
	const currentDir = "current_dir"

	data := []struct {
		name string
		mock interfaces.GetProjectName
	}{
		{
			name: currentDir,
			mock: mocks.IsCurrentDirIsTrue{},
		},
		{
			name: mocks.ProjName,
			mock: mocks.IsCurrentDirIsFalse{},
		},
	}

	for _, dataT := range data {
		{
			t.Log("create dir for ", dataT.name)
			{
				if err := os.Mkdir(dataT.name, 0755); err != nil {
					t.Fatalf("Failed to create test directory: %v", err)
				}

				if dataT.name == currentDir {
					if err := os.Chdir(fmt.Sprintf("./%s", currentDir)); err != nil {
						t.Fatalf("Expected no error, but got: %s", err)
					}
				}
			}

			t.Log("create repository in dir: ", dataT.name)
			{
				if err := CreateLocalGitRepository(dataT.mock); err != nil {
					t.Fatalf("Expected no error, but got: %s", err)
				}
			}

			t.Log("check the .git in the directory: ", dataT.name)
			{
				_, err := os.Stat(".git")
				if os.IsNotExist(err) {
					t.Fatal(".git directory not found in the expected location")
				}
			}

			if err := os.Chdir(".."); err != nil {
				t.Fatalf("Expected no error, but got: %s", err)
			}

			os.RemoveAll(fmt.Sprintf("./%s", dataT.name))

			t.Log("SUCCESS")
		}
	}
}
