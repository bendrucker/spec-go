package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() (*http.ServeMux, *httptest.Server, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	return mux, server, func() {
		server.Close()
	}
}

func TestHttpHead(t *testing.T) {
	resp, err := http.Head("https://golang.org")

	if !assert.NoError(t, err) {
		return
	}

	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
}

func TestHttptestMux(t *testing.T) {
	mux, server, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodHead, r.Method)
	})

	resp, err := http.Head(server.URL)

	if !assert.NoError(t, err) {
		return
	}

	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
}

func TestHttptestStatus(t *testing.T) {
	mux, server, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodHead, r.Method)
		w.WriteHeader(404)
	})

	resp, err := http.Head(server.URL)

	if !assert.NoError(t, err) {
		return
	}

	defer resp.Body.Close()
	assert.Equal(t, 404, resp.StatusCode)
}
