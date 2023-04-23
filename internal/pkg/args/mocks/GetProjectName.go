package mocks

const ProjName = "new"

type IsCurrentDirIsFalse struct{}

func (i IsCurrentDirIsFalse) GetProjectName() (string, bool, error) {
	return ProjName, false, nil
}

type IsCurrentDirIsTrue struct{}

func (i IsCurrentDirIsTrue) GetProjectName() (string, bool, error) {
	return "", true, nil
}
