package cmd

import (
	"fmt"
	"os"

	"github.com/eskog/projmanager/internal/helpers"
)

const GITHUBFOLDER = "/.github"
const DEPENDABOT = "assets/Dependabot/dependabot.yml.go"

func dependabot() error {
	dst, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Unable to get PWD: %w", err)
	}

	if !dirExists(dst + GITHUBFOLDER) {
		err = os.Mkdir(dst+GITHUBFOLDER, 0755)
		if err != nil {
			return fmt.Errorf("Unable to create directory at path: %s: %w", dst+GITHUBFOLDER, err)
		}
	}

	err = helpers.Filecopy(DEPENDABOT, dst+GITHUBFOLDER+"/dependabot.yml")
	if err != nil {
		return fmt.Errorf("Unable to copy dependabot.yml to destination: %w", err)
	}

	return nil
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && info.IsDir()
}
