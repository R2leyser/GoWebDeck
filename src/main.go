package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Script struct {
	ID          int    `json:"id"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Toggle 		bool   `json:"toggle"`
}

type Config struct {
	PrimaryForegroundColor   string `json:"primary-foreground"`
	PrimaryBackgroundColor   string `json:"primary-background"`
	SecondaryForegroundColor string `json:"secondary-foreground"`
	SecondaryBackgroundColor string `json:"secondary-background"`
}


var (
	scriptMap = make(map[int]Script)
	config    Config
	scripts   []Script
	nextID    = 1
)

func main() {
	http.HandleFunc("/", htmlHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	scriptsInit()

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	parseScripts(os.Getenv("HOME") + "/.config/gowebdeck/scripts.json")
	parseConfig(os.Getenv("HOME") + "/.config/gowebdeck/config.json")

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	tmplFile := "./static/index.tmpl"

	templateData := struct {
		Scripts                  []Script
		PrimaryForegroundColor   string
		PrimaryBackgroundColor   string
		SecondaryForegroundColor string
		SecondaryBackgroundColor string
	}{
		Scripts:                  scripts,
		PrimaryForegroundColor:   config.PrimaryForegroundColor,
		PrimaryBackgroundColor:   config.PrimaryBackgroundColor,
		SecondaryForegroundColor: config.SecondaryForegroundColor,
		SecondaryBackgroundColor: config.SecondaryBackgroundColor,
	}

	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, templateData)
	if err != nil {
		fmt.Println(err)
	}
}

func parseConfig(path string) {

	jsonFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(jsonFile), &config); err != nil {
		panic(err)
	}
}
