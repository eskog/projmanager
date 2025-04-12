package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const basePath = "/Users/skogen/Projects/Templates"

func Filecopy(src, dst string) error {
	if err := verifyPaths(src, dst); err != nil {
		return fmt.Errorf("filepaths incorrect: %w", err)
	}
	content, err := os.ReadFile(filepath.Clean(src))
	if err != nil {
		return fmt.Errorf("unable to read template file %s: %w", src, err)
	}
	file, err := os.Create(filepath.Clean(dst))
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

func verifyPaths(src, dst string) error {
	if !strings.HasPrefix(src, basePath) {
		return fmt.Errorf("invalid path to template file: %s", src)
	}

	workingdir, _ := os.Getwd()
	if !strings.HasPrefix(dst, workingdir) {
		return fmt.Errorf("invalid destination path: %s", dst)
	}

	return nil
}
