package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

var PORT = os.Getenv("PORT")

const FSPATH = "./static/"
const IndexHTML = "/index.html"
const CSS = "/temp.css"
const IMAGE = "/gear.png"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func logIncomingRequest(r *http.Request) {
	log.WithFields(log.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"req":    r.TLS,
		"proto":  r.Proto,
		"host":   r.Host,
		"remote": r.RemoteAddr,
	}).Info("Request received for static assets")
}

func logResponseToFileNotFound(r *http.Request, fullPath string, altPath string) {
	log.WithFields(log.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"req":    r.TLS,
		"proto":  r.Proto,
		"host":   r.Host,
		"remote": r.RemoteAddr,
	}).Info(fmt.Sprintf("Requested %s does not exist, returning %s file", fullPath, altPath))
}

func handleFileNotFound(fullPath string, r *http.Request) {
	if strings.Contains(fullPath, CSS) {
		logResponseToFileNotFound(r, fullPath, CSS)
		r.URL.Path = CSS
	} else if strings.Contains(fullPath, IMAGE) {
		logResponseToFileNotFound(r, fullPath, IMAGE)
		r.URL.Path = IMAGE
	} else {
		logResponseToFileNotFound(r, fullPath, IndexHTML)
		r.URL.Path = "/"
	}
}

func serveStaticAssets(w http.ResponseWriter, r *http.Request) {
	logIncomingRequest(r)

	fs := http.FileServer(http.Dir(FSPATH))

	requestUri := r.URL.Path

	if requestUri != "/" {
		_, err := os.Stat(requestUri)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(err)
			}
			handleFileNotFound(requestUri, r)
		}
	}

	fs.ServeHTTP(w, r)
}

func startServer(server *http.Server) {
	log.Info("Starting server on PORT " + PORT)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", serveStaticAssets)

	Addr := fmt.Sprintf(":%s", PORT)

	server := &http.Server{
		Addr: Addr,
	}

	startServer(server)
}
