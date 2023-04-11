package version

import (
	"os/exec"
	"regexp"
	"strings"
)

// GoVersion receives the current version of go and returns it in go X.Y format.
func GoVersion() (string, error) {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// to search for the desired part in the string.
	re := regexp.MustCompile(`go\d+\.\d+(\.\d+)?`)
	version := re.FindString(string(output))

	// convert to the desired form.
	version = strings.Replace(version, "go", "", 1)
	versionParts := strings.Split(version, ".")
	version = "go " + versionParts[1] + "." + versionParts[2]

	return version, nil
}
