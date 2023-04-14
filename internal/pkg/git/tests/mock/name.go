package mock

type IsCurrentDirIsFalse struct{}

const TestDir = "test_dir"

func (i IsCurrentDirIsFalse) GetProjectName() (string, bool, error) {
	return TestDir, false, nil
}

type IsCurrentDirIsTrue struct{}

func (i IsCurrentDirIsTrue) GetProjectName() (string, bool, error) {
	return "", true, nil
}
