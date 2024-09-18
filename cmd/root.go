package cmd

import (
	"fmt"
	"movielibrary/config"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:    "app",
	PreRun: preRun,
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func preRun(cmd *cobra.Command, args []string) {
	// init logger here!!! or in main.go
	config.Load()
}

// Execute executes the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
