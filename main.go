package main

import (
	"fmt"
	"html"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
