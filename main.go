package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func init() {
	log.SetFormatter(&log.JSONFormatter{})
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
	http.ServeFile(w, r, "static/"+r.URL.Path)
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
	Addr := fmt.Sprintf(":%s", port)
	server := &http.Server{
		Addr: Addr,
	}

	startServer(server)
}
