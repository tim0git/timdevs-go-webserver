package main

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestReturnsStatus200(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	staticWebServer(w, req)

	res := w.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestReturnsHTML(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	staticWebServer(w, req)

	res := w.Result()

	assert.Contains(t, res.Header.Get("Content-Type"), "text/html")
}

func TestReturnsHTMLOnAnyPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown/", nil)
	w := httptest.NewRecorder()

	staticWebServer(w, req)

	res := w.Result()

	assert.Contains(t, res.Header.Get("Content-Type"), "text/html")
}

func TestLogsRequestsToStdOut(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown/", nil)
	w := httptest.NewRecorder()

	stdOut := captureOutput(func() {
		staticWebServer(w, req)
	})

	assert.Contains(t, stdOut, "\"level\":\"info\"")
}
