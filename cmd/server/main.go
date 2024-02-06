package main

import (
	"flag"
	"fmt"
	"go/build"
	"net/http"
	"path/filepath"
	"text/template"
)

type Page struct {
	Title string
}

func serve(addr string) error {

	pagedata := Page{
		Title: "osiris web server",
	}

	homeTemplate := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		url := r.URL.Path

		fmt.Println("url", url)

		if url == "/wasm_exec.js" {
			http.FileServer(http.Dir(filepath.Join(build.Default.GOROOT, "misc/wasm/"))).ServeHTTP(w, r)
			return
		}

		if url == "/demo.wasm" {
			http.FileServer(http.Dir("/usr/local/src/cmd/client")).ServeHTTP(w, r)
			return
		}

		err := homeTemplate.Execute(w, pagedata)
		if err != nil {
			panic(err)
		}
	})

	// Start the server
	// nolint: gosec
	return http.ListenAndServe(addr, nil)
}

func main() {
	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()
	fmt.Println("Server is running on", *addr)
	err := serve(*addr)
	if err != nil {
		panic(err)
	}

}
