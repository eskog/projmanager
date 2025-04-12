/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := gitignore(); err != nil {
			log.Printf("Error creating gitignore: %s", err)
		}
		if err := dependabot(); err != nil {
			log.Printf("error creating dependabot.yml: %s", err)
		}
	},
}

var initGitIgnoreCmd = &cobra.Command{
	Use:   "gitignore",
	Short: "init a .gitignore file",
	Long:  `Initialize an default .gitignore file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := gitignore(); err != nil {
			log.Fatal(err)
		}
	},
}

var initDependabotCmd = &cobra.Command{
	Use:   "dependabot",
	Short: "init a dependabot.yml file",
	Long:  `Initialize an default dependabot.yml file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := dependabot(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.AddCommand(initGitIgnoreCmd, initDependabotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
