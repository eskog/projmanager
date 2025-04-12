package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/eskog/projmanager/internal/helpers"
)

const GOGITIGNORE = "/Users/skogen/Projects/Templates/gitignore/go.gitignore"

func gitignore() error {

	dst, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Unable to get PWD: %w", err)
	}
	dst = dst + "/.gitignore"
	if err = helpers.Filecopy(GOGITIGNORE, dst); err != nil {
		return fmt.Errorf("unable to add .gitignore file: %w", err)
	}
	log.Println(".gitignore added successfully")
	return nil
}
