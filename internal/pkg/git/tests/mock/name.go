package mock

const TestDir = "test_dir"

type IsCurrentDirIsFalse struct{}

func (i IsCurrentDirIsFalse) GetProjectName() (string, bool, error) {
	return TestDir, false, nil
}

type IsCurrentDirIsTrue struct{}

func (i IsCurrentDirIsTrue) GetProjectName() (string, bool, error) {
	return "", true, nil
}
