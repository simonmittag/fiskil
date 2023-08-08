package integration

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

func TestIntegration(t *testing.T) {
	url := "http://localhost:8080"
	r := 100

	var wg sync.WaitGroup
	wg.Add(r)

	for i := 0; i < r; i++ {
		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()
		}()
	}

	wg.Wait()
	fmt.Println("All requests completed.")
}
