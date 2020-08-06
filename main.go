package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/index.html"))
	err := t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./download")
	if err != nil {
		log.Fatal(err)
	}

	folder := []string{}

	for _, f := range files {
		if f.IsDir() {
			folder = append(folder, f.Name())
		}
	}

	json.NewEncoder(w).Encode(folder)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	category := vars["category"][0]

	files, err := ioutil.ReadDir("./download/" + category)
	if err != nil {
		log.Fatal(err)
	}

	filez := []string{}

	for _, f := range files {
		if !f.IsDir() {
			filez = append(filez, f.Name())
		}
	}

	json.NewEncoder(w).Encode(filez)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir("./download"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/categories", categoriesHandler)
	http.HandleFunc("/file", fileHandler)
	http.ListenAndServe(":8888", nil)
}
