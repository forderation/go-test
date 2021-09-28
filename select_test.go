package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var tendSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tendSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-Ping(a):
		return a, nil
	case <-Ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func Ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func MeasureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"

	want := fastURL
	got, _ := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestServerRace(t *testing.T) {
	slowServer := MakeDelayedServer(12 * time.Millisecond)
	fastServer := MakeDelayedServer(11 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()
	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got, _ := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func MakeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			rw.WriteHeader(http.StatusOK)
		},
	))
}

func TestCompiledServerRace(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := MakeDelayedServer(20 * time.Millisecond)
		fastServer := MakeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := MakeDelayedServer(25 * time.Millisecond)
		defer server.Close()
		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
