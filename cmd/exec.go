/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/leonardogregoriocs/load_tester_cli/internal/report"
	"github.com/leonardogregoriocs/load_tester_cli/internal/tester"
	"github.com/leonardogregoriocs/load_tester_cli/utils"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
func ExecuteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Load Tester CLI em Go",
		Short: "This command executes the flow",
		Long:  `Providing the service URL, the total number of desired requests, and the quantity of simultaneous calls`,
		Run: func(cmd *cobra.Command, args []string) {
			url, _ := cmd.Flags().GetString("url")
			request, _ := cmd.Flags().GetString("requests")
			concurrency, _ := cmd.Flags().GetString("concurrency")

			ok := utils.CheckFlags(url, request, concurrency)
			if !ok {
				cmd.Help()
			}

			result := tester.ExecuteTest(url, utils.ConvertNumber(request), utils.ConvertNumber(concurrency))

			report.GenerateReport(result)
		},
	}

	cmd.Flags().StringP("url", "u", "", "URL of the service to be tested")
	cmd.Flags().StringP("requests", "r", "", "total number of requests")
	cmd.Flags().StringP("concurrency", "c", "", "number of simultaneous calls")

	return cmd
}
