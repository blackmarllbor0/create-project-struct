package args

import (
	"errors"
	"os"
)

// GetProjectName функция для получения имени проекта.
// Имя задается через аргумент командной строки.
// Задайте аргумент ".", чтобы создать проект в текущей директории.
// Булево значение возвращает true в том случае, если приложение
// создается в текущем каталоге.
func GetProjectName() (string, bool, error) {
	projectName := os.Args[1]

	if projectName == "" {
		return "", false, errors.New("имя проекта не было передано")
	}

	if projectName == "." {
		pwd, err := os.Getwd()
		if err != nil {
			return "", false, err
		}
		return pwd, true, nil
	}

	return projectName, false, nil
}
