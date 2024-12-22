package request_id

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestIdPlugin(t *testing.T) {
	cfg := CreateConfig()
	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "request-id-plugin")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Adds X-Request-Id if missing", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
		if err != nil {
			t.Fatal(err)
		}

		handler.ServeHTTP(recorder, req)

		assertHeaderExists(t, req, "X-Request-Id")
	})

	t.Run("Does not overwrite existing X-Request-Id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
		if err != nil {
			t.Fatal(err)
		}

		existingID := "existing-id"
		req.Header.Set("X-Request-Id", existingID)

		handler.ServeHTTP(recorder, req)

		assertHeaderValue(t, req, "X-Request-Id", existingID)
	})
}

func assertHeaderExists(t *testing.T, req *http.Request, key string) {
	t.Helper()

	value := req.Header.Get(key)
	if value == "" {
		t.Errorf("expected header %s to exist, but it was not set", key)
	}
}

func assertHeaderValue(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	value := req.Header.Get(key)
	if value != expected {
		t.Errorf("invalid header value for %s: expected %s, got %s", key, expected, value)
	}
}
