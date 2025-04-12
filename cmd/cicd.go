package cmd

import (
	"fmt"
	"os"

	"github.com/eskog/projmanager/internal/helpers"
)

const WORKFLOWS = "/.github/workflows"
const CI = "assets/Workflow/CI.yml"
const CD = "assets/Workflow/CD.yml"

func cicd() error {
	dst, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Unable to get PWD: %s", err)
	}

	if !helpers.DirExists(dst + WORKFLOWS) {
		if err := os.MkdirAll(dst+WORKFLOWS, 0755); err != nil {
			return fmt.Errorf("Unable to create dir %s: %w", dst+WORKFLOWS, err)
		}
	}

	if err := helpers.Filecopy(CI, dst+WORKFLOWS+"/CI.yml"); err != nil {
		return fmt.Errorf("Unable to copy CI.yml to dst %s: %w", dst+WORKFLOWS+"/CI.yml", err)
	}
	if err := helpers.Filecopy(CI, dst+WORKFLOWS+"/CD.yml"); err != nil {
		return fmt.Errorf("Unable to copy CD.yml to dst %s: %w", dst+WORKFLOWS+"/CD.yml", err)
	}

	return nil
}
