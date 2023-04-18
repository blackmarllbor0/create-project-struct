package temp

import (
	"errors"
	"io/fs"
	"log"
	"path/filepath"
)

const (
	MAIN = "main.go"
	MAKE = "make"
	LINT = "lint"
)

// TODO we need to improve performance and fix minor errors.

func GetTemplateByAlias(alias string) (string, error) {
	var pathToTemp string

	err := filepath.WalkDir("../../../", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == alias {
			files, err := filepath.Glob(filepath.Join(path, "*"))
			if err != nil {
				return err
			}

			for _, fl := range files {
				if filepath.Base(fl) == alias {
					log.Println(filepath.Base(fl))
					pathToTemp = fl
				} else {
					return errors.New("template with such alias does not exist")
				}

				break
			}
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return pathToTemp, err
}
