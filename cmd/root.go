package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/coolblknerd/go-init/cmd/project"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "go-init",
	Short: "Go-init is a golang project generator.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go-Init project generator")
	},
}

var createCmd = &cobra.Command{
	Use:   "create [project name]",
	Short: "Creates a new golang project in your workspace.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		proj := project.New(strings.Join(args, ""))
		proj.Create()
	},
}
