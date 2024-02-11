package tester

import (
	"fmt"
	"net/http"
	"time"
)

type TestResult struct {
	TotalTime     time.Duration
	TotalRequests int
	Status200     int
	StatusOthers  map[int]int
}

func ExecuteTest(url string, totalRequests, concurrency int) TestResult {
	start := time.Now()
	reqChannel := make(chan int, concurrency)
	result := TestResult{StatusOthers: make(map[int]int)}

	for i := 0; i < totalRequests; i++ {
		reqChannel <- 1
		go func() {
			defer func() { <-reqChannel }()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error while making request:", err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				result.Status200++
			} else {
				result.StatusOthers[resp.StatusCode]++
			}
		}()
	}

	for i := 0; i < concurrency; i++ {
		reqChannel <- 1
	}

	result.TotalTime = time.Since(start)
	result.TotalRequests = totalRequests
	return result
}
