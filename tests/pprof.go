package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/liu602348184/go-captcha/tests/pprof"
)

func main() {
	// Example: demo
	http.HandleFunc("/handler", pprof.Handler)
	http.ListenAndServe(":9999", nil)
}
