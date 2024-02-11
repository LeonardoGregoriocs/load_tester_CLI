package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "load",
	Short: "This is a CLI system developed in Go to conduct load testing on web services",
	Long: `
	This is a CLI system developed in Go to conduct load testing on web services. 
	With this tool, you can test the performance of any HTTP service by providing the service URL, 
	the total number of desired requests, and the concurrency level.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(ExecuteCmd())
}
