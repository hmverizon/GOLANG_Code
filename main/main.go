package main

import (
	"net/http"
)

var test = "testing"

func main() {
	http.HandleFunc("/contact", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		arg1.Write([]byte("Hi hi hi"))
	})
	http.ListenAndServe("8090", nil)
}
