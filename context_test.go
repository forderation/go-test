package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	support "github.com/forderation/go-test/support"
)

func TestContextServer(t *testing.T) {
	data := "hello, world"

	t.Run("response normal case string", func(t *testing.T) {
		svr := support.Server(&support.StubStore{Response: data})
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &support.SpyStore{Response: data, T: t}
		svr := support.Server(store)

		timeout := 5 * time.Millisecond
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		// cancellingCtx, cancel := context.WithCancel(request.Context())
		// time.AfterFunc(timeout, cancel)
		// request = request.WithContext(cancellingCtx)

		// alternate using WithTimeout
		timeoutCtx, cancel := context.WithTimeout(request.Context(), timeout)
		defer cancel()
		request = request.WithContext(timeoutCtx)

		response := &support.SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.Written {
			t.Error("a response should not have been written")
		}
		store.AssertWasCancelled()
	})

	t.Run("returns data from store", func(t *testing.T) {
		store := &support.SpyStore{Response: data, T: t}
		svr := support.Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		store.AssertWasNotCancelled()
	})
}
