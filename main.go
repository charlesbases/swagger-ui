package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

const dist = "/swagger/"

var (
	// port port
	port = ":8888"
	// folder api folder
	folder = "./swagger/api"
)

func init() {
	if envPort := os.Getenv("SWAGGER_PORT"); len(envPort) != 0 {
		port = envPort
	}
	if envPath := os.Getenv("SWAGGER_PATH"); len(envPath) != 0 {
		folder = envPath
	}
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix(dist).Handler(fileserver())

	n := negroni.New()
	n.Use(recovery())
	n.UseHandler(r)

	n.Run(port)
}

// rec .
type rec struct{}

// recovery .
func recovery() *rec {
	return new(rec)
}

// ServeHTTP .
func (rec) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data, _ := json.Marshal(err)
			rw.Write(data)
		}
	}()

	next(rw, r)
}

type fileServer struct {
	handler http.Handler
	fsystem http.FileSystem
}

// fileserver .
func fileserver() *fileServer {
	return &fileServer{
		handler: http.FileServer(http.Dir(".")),
		fsystem: http.Dir(folder),
	}
}

// ServeHTTP .
func (fs *fileServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, ".json") {
		subpath := strings.TrimPrefix(r.URL.Path, dist)

		f, err := fs.fsystem.Open(subpath)
		if err != nil {
			http.Error(rw, fmt.Sprintf("%s not found or permission denied", subpath), http.StatusNotFound)
			return
		}

		var buf = new(bytes.Buffer)
		io.Copy(buf, f)
		f.Close()

		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		rw.WriteHeader(http.StatusOK)
		rw.Write(buf.Bytes())
		return
	}

	fs.handler.ServeHTTP(rw, r)
}
