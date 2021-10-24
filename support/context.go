package support

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}

type StubStore struct {
	Response  string
	Cancelled bool
}

func (s *StubStore) Fetch(ctx context.Context) (string, error) {
	return s.Response, nil
}

func (s *StubStore) Cancel() {
	s.Cancelled = true
}

type SpyStore struct {
	Response  string
	Cancelled bool
	T         *testing.T
}

func (s *SpyStore) AssertWasCancelled() {
	s.T.Helper()
	if !s.Cancelled {
		s.T.Errorf("store was not told to cancel")
	}
}

func (s *SpyStore) AssertWasNotCancelled() {
	s.T.Helper()
	if s.Cancelled {
		s.T.Error("store was told to cancel")
	}
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.Response {
			select {
			case <-ctx.Done():
				s.T.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()
	select {
	case <-ctx.Done():
		s.Cancelled = true
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {
	s.Cancelled = true
}

type SpyResponseWriter struct {
	Written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.Written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.Written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.Written = true
}
