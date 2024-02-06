package main

import (
	"flag"
	"net/http"
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

	err := serve(*addr)
	if err != nil {
		panic(err)
	}
}
