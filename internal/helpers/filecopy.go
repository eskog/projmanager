package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const basePath = "/Users/skogen/Projects/Templates"

func Filecopy(src, dst string) error {
	src, dst, err := verifypaths(src, dst)
	if err != nil {
		return fmt.Errorf("error sanitizing filepaths: %w", err)
	}
	content, err := os.ReadFile(src) // #nosec G304 Due to false alarm, the paths are sanitized in the call to verifypaths
	if err != nil {
		return fmt.Errorf("unable to read template file %s: %w", src, err)
	}
	file, err := os.Create(dst) // #nosec G304 Due to false alarm, the paths are sanitized in the call to verifypaths
	if err != nil {
		return fmt.Errorf("unable to create empty destination file:%s: %w", dst, err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("unable to write to %s: %w", dst, err)
	}

	return nil
}

func verifypaths(src, dst string) (string, string, error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)
	if !strings.HasPrefix(src, basePath) {
		return "", "", fmt.Errorf("invalid path to template file: %s", src)
	}

	workingdir, _ := os.Getwd()
	if !strings.HasPrefix(dst, workingdir) {
		return "", "", fmt.Errorf("invalid destination path: %s", dst)
	}

	return src, dst, nil
}
