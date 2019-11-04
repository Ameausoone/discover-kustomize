package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//m := make(map[string]string)
	log.Printf("youhou")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			//m[e[:i]] = e[i+1:]
			key := e[:i]
			log.Printf("Key: %s, value: %s\n", e[:i], e[i+1:])
			if strings.HasPrefix(key, "app_") {
				fmt.Fprintf(w, "<h1>Key: %s, Value: %s</h1>", e[:i], e[i+1:])
			}
		}
	}

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
