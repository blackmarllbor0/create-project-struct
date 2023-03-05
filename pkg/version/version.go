package version

import (
	"os/exec"
	"regexp"
	"strings"
)

// GoVersion получает текущую версию go и возвращает ее в формате go X.Y.
func GoVersion() (string, error) {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// для поиска нужной части в строке.
	re := regexp.MustCompile(`go\d+\.\d+(\.\d+)?`)
	version := re.FindString(string(output))

	// преобразуем к нужному виду.
	version = strings.Replace(version, "go", "", 1)
	versionParts := strings.Split(version, ".")
	version = "go " + versionParts[1] + "." + versionParts[2]

	return version, nil
}
