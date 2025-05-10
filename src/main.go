package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "html/template"
	"sync"
)

type Script struct {
	ID          int    `json:"id"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

var (
	scriptMap = make(map[int]Script)
	scripts []Script
	nextID    = 1
	postsMu   sync.Mutex
)

func main() {
	http.HandleFunc("/", htmlHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	scriptsInit();

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	parseScripts(os.Getenv("HOME") + "/.config/gowebdeck/scripts.json")

    w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(http.StatusOK)

	tmplFile := "./static/index.tmpl"

	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, scripts)
	if err != nil {
		fmt.Println(err)
	}
}

