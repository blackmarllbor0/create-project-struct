package args

import (
	"os"
	"testing"
)

func TestGetProjectName(t *testing.T) {
	projName := ProjectName{}

	t.Log("no project name provided")
	{
		projectName, _, err := projName.GetProjectName()
		if err != nil {
			t.Fatalf("Expected an error, got nil")
		}

		if projectName == "" {
			t.Fatalf("Expected empty project name, got '%s'", projectName)
		}
	}

	t.Log("projectName is \".\"")
	{
		os.Args = []string{"cmd", "."}
		projectName, isCurrentDir, err := projName.GetProjectName()
		if err != nil {
			t.Fatalf("Expected an error, got nil")
		}

		pwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("Fail to get current directory: %s", err.Error())
		}

		if projectName != pwd {
			t.Fatalf("Expected project name to be: %s, got: %s", pwd, projectName)
		}

		if !isCurrentDir {
			t.Fatalf("Expected isCurrentDir to be true, got false")
		}
	}

	t.Log("project name is a valid directory path")
	{
		testProjectName := "test_project"
		os.Args = []string{"cmd", testProjectName}
		projectName, isCurrentDir, err := projName.GetProjectName()
		if err != nil {
			t.Fatalf("Expected an error, got nil")
		}

		if projectName != testProjectName {
			t.Fatalf("Expected project name to be '%s', got '%s'", testProjectName, projectName)
		}

		if isCurrentDir {
			t.Fatalf("Expected isCurrentDir to be false, got true")
		}
	}

	t.Log("SUCCESS")
}
