package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/eskog/projmanager/internal/helpers"
)

const GITHUBFOLDER = "/.github"
const DEPENDABOT = "/Users/skogen/Projects/Templates/Dependabot/dependabot.yml.go"

func dependabot() error {
	dst, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get PWD: %w", err)
	}

	if !helpers.DirExists(dst + GITHUBFOLDER) {
		err = os.Mkdir(dst+GITHUBFOLDER, 0750)
		if err != nil {
			return fmt.Errorf("unable to create directory at path: %s: %w", dst+GITHUBFOLDER, err)
		}
	}

	err = helpers.Filecopy(DEPENDABOT, dst+GITHUBFOLDER+"/dependabot.yml")
	if err != nil {
		return fmt.Errorf("unable to copy dependabot.yml to destination: %w", err)
	}
	log.Println("dependabot.yml added successfully")

	return nil
}
