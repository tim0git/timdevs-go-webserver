package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

var PORT = os.Getenv("PORT")

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Starting server on PORT " + PORT)

	log.Fatal(srv.ListenAndServe())
}
