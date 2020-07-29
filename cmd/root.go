package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// func init() {
// 	rootCmd.AddCommand(versionCmd)
// }

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
