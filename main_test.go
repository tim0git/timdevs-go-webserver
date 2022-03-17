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

func TestReturnsStatus200OnAnUnknownPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown/", nil)
	w := httptest.NewRecorder()

	staticWebServer(w, req)

	res := w.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestReturnsContentTypeHTMLOnRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	staticWebServer(w, req)

	res := w.Result()

	assert.Equal(t, "text/html; charset=utf-8", res.Header.Get("Content-Type"))
}

func TestReturnsContentTypeCSSOnRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/temp.css", nil)
	w := httptest.NewRecorder()

	staticWebServer(w, req)

	res := w.Result()

	assert.Equal(t, "text/css; charset=utf-8", res.Header.Get("Content-Type"))
}

func TestReturnsContentTypeGearPNGOnRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/gear.png", nil)
	w := httptest.NewRecorder()

	staticWebServer(w, req)

	res := w.Result()

	assert.Equal(t, "image/png", res.Header.Get("Content-Type"))
}

func TestLogsRequestsToStdOut(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	stdOut := captureOutput(func() {
		staticWebServer(w, req)
	})

	assert.Contains(t, stdOut, "\"level\":\"info\"")
}

func TestReturnsHTMLOnAnyPathOtherThanRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/sainsburysbank/", nil)
	w := httptest.NewRecorder()

	staticWebServer(w, req)

	res := w.Result()

	assert.Equal(t, "text/html; charset=utf-8", res.Header.Get("Content-Type"))
}

func TestReturnsContentTypeCSSOnAnyPathOtherThanRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown/temp.css", nil)
	w := httptest.NewRecorder()

	staticWebServer(w, req)

	res := w.Result()

	assert.Equal(t, "text/css; charset=utf-8", res.Header.Get("Content-Type"))
}

func TestReturnsContentTypeGearPNGOnAnyPathOtherRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown/gear.png", nil)
	w := httptest.NewRecorder()

	staticWebServer(w, req)

	res := w.Result()

	assert.Equal(t, "image/png", res.Header.Get("Content-Type"))
}
