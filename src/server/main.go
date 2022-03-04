package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var port = ":8080"

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

func logger(r *http.Request) {
	log.WithFields(log.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"req":    r.TLS,
		"proto":  r.Proto,
		"host":   r.Host,
		"remote": r.RemoteAddr,
	}).Info("Request received for static assets")
}

func staticWebServer(w http.ResponseWriter, r *http.Request) {
	logger(r)
	http.ServeFile(w, r, "./static/index.html")
}

func startServer(server *http.Server) {
	log.Info("Starting server on port " + port)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", staticWebServer)

	server := &http.Server{
		Addr: port,
	}

	startServer(server)
}
