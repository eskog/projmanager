package cmd

import (
	"fmt"
	"os"

	"github.com/eskog/projmanager/internal/helpers"
)

const WORKFLOWS = "/.github/workflows"
const CI = "/Users/skogen/Projects/Templates/Workflow/CI.yml"
const CD = "/Users/skogen/Projects/Templates/Workflow/CD.yml"

func cicd() error {
	dst, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get PWD: %s", err)
	}

	if !helpers.DirExists(dst + WORKFLOWS) {
		if err := os.MkdirAll(dst+WORKFLOWS, 0750); err != nil {
			return fmt.Errorf("unable to create dir %s: %w", dst+WORKFLOWS, err)
		}
	}

	if err := helpers.Filecopy(CI, dst+WORKFLOWS+"/CI.yml"); err != nil {
		return fmt.Errorf("unable to copy CI.yml to dst %s: %w", dst+WORKFLOWS+"/CI.yml", err)
	}
	if err := helpers.Filecopy(CI, dst+WORKFLOWS+"/CD.yml"); err != nil {
		return fmt.Errorf("unable to copy CD.yml to dst %s: %w", dst+WORKFLOWS+"/CD.yml", err)
	}

	return nil
}
