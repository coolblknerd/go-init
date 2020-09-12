package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/coolblknerd/go-init/run/create"
	"github.com/coolblknerd/go-init/run/create/api"
	"github.com/spf13/cobra"
)

var (
	pkgName string
	handles []string
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createAPI)
	createAPI.Flags().StringVar(&pkgName, "name", "", "Name of package")
	createAPI.Flags().StringArrayVar(&handles, "handles", handles, "Name of handlers for API")
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
		proj := create.New(strings.Join(args, ""))
		proj.Create()
	},
}

var createAPI = &cobra.Command{
	Use:   "api [api name]",
	Short: "Creates a new golang API project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName, _ = cmd.Flags().GetString("name")
		handles, _ = cmd.Flags().GetStringArray("handles")

		proj := api.New(pkgName, handles)
		proj.Create()
	},
}
