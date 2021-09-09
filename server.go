package main

import (
	"flag"
	"log"
	"net/http"
    "strings"
)

var fileServer = http.FileServer(http.Dir("."))
var mimeList = [][]string{{".js", "text/javascript; charset=UTF-8"}}

func fileHandler(w http.ResponseWriter, r *http.Request) {
        ruri := r.RequestURI
        for _, s := range mimeList {
            if strings.HasSuffix(ruri, s[0]) {
                    w.Header().Set("Content-Type", s[1])
            }
        }

        fileServer.ServeHTTP(w, r)
        log.Printf("- [%s %s] [%s]", r.Method, r.RequestURI, r.UserAgent())

}

func main() {
	port := flag.String("p", "8080", "port")
	flag.Parse()

	http.HandleFunc("/", fileHandler)

	log.Printf("- Listening on: %s", *port)
	http.ListenAndServe(":"+*port, nil)
}