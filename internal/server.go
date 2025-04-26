package internal

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gomarkdown/markdown"
	"io"
	"net/http"
	"strings"
)

type serverT struct {
	config Config
}

var Server serverT

// ListenAndServe starts the web server
func (serverT) ListenAndServe() {
	Server.config = ReadConfig()
	docsDir := fmt.Sprintf("./%s", Server.config.Directories.Docs)
	publicDir := fmt.Sprintf("./%s", Server.config.Directories.Public)

	fs := http.Dir(docsDir)
	fsPublic := http.FileServer(http.Dir(publicDir))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// handle routes
		if targetFile, ok := Server.config.Routes[r.URL.Path]; ok {
			file, err := fs.Open(targetFile)
			if err != nil {
				http.NotFound(w, r)
				return
			}
			defer func(file http.File) {
				err := file.Close()
				if err != nil {
					log.Error("failed to open target file", err)
				}
			}(file)

			mdContent, err := io.ReadAll(file)
			if err != nil {
				http.Error(w, "Error reading file", http.StatusInternalServerError)
				log.Error("failed to read file", err)
				return
			}

			html := markdown.ToHTML(mdContent, nil, nil)
			w.Header().Set("Content-Type", "text/html")
			_, err = w.Write(html)
			if err != nil {
				log.Error("failed to write data to socket", err)
			}
			return
		}

		// handle direct .md file requests
		if strings.HasSuffix(r.URL.Path, ".md") {
			file, err := fs.Open(r.URL.Path)
			if err != nil {
				http.NotFound(w, r)
				return
			}
			defer func(file http.File) {
				err := file.Close()
				if err != nil {
					log.Error("failed to close file", err)
				}
			}(file)

			mdContent, err := io.ReadAll(file)
			if err != nil {
				http.Error(w, "Error reading file", http.StatusInternalServerError)
				log.Error("failed to read file", err)
				return
			}

			html := markdown.ToHTML(mdContent, nil, nil)
			w.Header().Set("Content-Type", "text/html")
			_, err = w.Write(html)
			if err != nil {
				log.Error("failed to write data to socket", err)
			}
			return
		}

		// serve static files from public directory
		fsPublic.ServeHTTP(w, r)
	})

	log.Info("starting the http server")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", Server.config.Server.Host, Server.config.Server.Port), nil)
	if err != nil {
		log.Error("failure whilst starting the http server", err)
	}
}
