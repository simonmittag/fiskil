package integration

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"testing"
)

// This integration test is meant to be run against a live server that has just started and not processed any requests
func TestIntegration(t *testing.T) {
	url := "http://localhost:8080"
	n := 10
	//make 10 req in parallel
	makeNRequests(url, 10)

	//now GET #11
	resp, e := http.Get(url)
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		t.Errorf("request error: %v", e)
	}
	b := string(body)
	if !strings.Contains(b, strconv.Itoa(n+1)) {
		t.Errorf("invalid sequence number, got %v", b)
	} else {
		t.Logf("contains expected sequence number %v", b)
	}
}

func makeNRequests(url string, n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, _ := http.Get(url)
			defer resp.Body.Close()
		}()
	}
	wg.Wait()
}
