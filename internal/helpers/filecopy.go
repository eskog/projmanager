package helpers

import (
	"fmt"
	"os"
)

func Filecopy(src, dst string) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("Unable to read template file %s: %w", src, err)
	}
	file, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Unable to create empty destination file:%s: %w", dst, err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("Unable to write to %s: %w", dst, err)
	}

	return nil
}
