package api

import (
	"fmt"
	"net/http"
)

func PrintHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
