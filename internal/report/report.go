package report

import (
	"fmt"

	"github.com/leonardogregoriocs/load_tester_cli/internal/tester"
)

func GenerateReport(result tester.TestResult) {
	fmt.Println("Load Test Report:")
	fmt.Printf("Total Time Spent: %v\n", result.TotalTime)
	fmt.Printf("Total Number of Requests: %d\n", result.TotalRequests)
	fmt.Printf("Requests with Status 200: %d\n", result.Status200)
	fmt.Println("Distribution of Other Status Codes:")
	for code, count := range result.StatusOthers {
		fmt.Printf("Status %d: %d\n", code, count)
	}
}
