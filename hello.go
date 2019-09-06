// +build ignore

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chai2010/lena-go"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, lena.Data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
