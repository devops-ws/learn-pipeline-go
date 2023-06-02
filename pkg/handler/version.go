package handler

import (
	"fmt"
	"log"
	"net/http"
)

var version string

type Version struct {
	Log *log.Logger
}

func (h *Version) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ \"version\": \"1.0.0\"}")
}
