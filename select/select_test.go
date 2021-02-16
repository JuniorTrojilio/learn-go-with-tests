package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	SlowServer := createDelayServer(1 * time.Millisecond)
	FastServer := createDelayServer(0 * time.Millisecond)

	defer SlowServer.Close()
	defer FastServer.Close()

	SlowURL := SlowServer.URL
	FastURL := FastServer.URL

	expect := FastURL
	result, err := Runner(SlowURL, FastURL)

	if err != nil {
		t.Fatalf("dont expect a error, but have %v", err)
	}

	if result != expect {
		t.Errorf("result '%s', expect '%s'", result, expect)
	}

	t.Run("return error if the server not response in 10 seconds", func(t *testing.T) {
		server := createDelayServer(25 * time.Millisecond)

		defer server.Close()

		_, err := Config(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("Expected an error but it didn't happen")
		}
	})
}

func createDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
