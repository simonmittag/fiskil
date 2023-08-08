package integration

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestIntegration(t *testing.T) {
	url := "http://localhost:8080"
	r := 10

	var wg sync.WaitGroup

	for i := 0; i < r; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, _ := http.Get(url)
			defer resp.Body.Close()
		}()
	}

	wg.Wait()

	resp, e := http.Get(url)
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		t.Errorf("request error: %v", e)
	}
	b := string(body)
	if !strings.Contains(b, strconv.Itoa(r+1)) {
		t.Errorf("invalid sequence number, got %v", b)
	} else {
		t.Logf("expected sequence %v", b)
	}
}
