package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
    "html/template"
	"strconv"
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
	http.HandleFunc("/scripts/", scriptHandler)
	http.HandleFunc("/", htmlHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Script added:", scriptMap[nextID-1].Path)

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

func scriptHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/scripts/"):])
	if err != nil {
		http.Error(w, "Invalid Script ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		handlePostScript(w, r, id)
	case "GET":
		handleGetScript(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePostScript(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	p, ok := scriptMap[id]
	if !ok {
		http.Error(w, "Script not found", http.StatusNotFound)
		return
	}

	go executeScript(p.Path)
	http.Error(w, "Script execution started", http.StatusAccepted)
}

func handleGetScript(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	// If you use a two-value assignment for accessing a
	// value on a map, you get the value first then an
	// "exists" variable.
	_, ok := scriptMap[id]
	if !ok {
		http.Error(w, "Script not found", http.StatusNotFound)
		return
	}

	delete(scriptMap, id)
	w.WriteHeader(http.StatusOK)
}

func executeScript(path string) {
	// Simulate script execution
	// Here you would add the actual script execution logic

	pathFound, err := exec.LookPath(path)
	if err != nil {
		fmt.Println("Error finding script:", err)
		return
	}
	cmd := exec.Command(pathFound)

	error := cmd.Run()
	if error != nil {
		fmt.Println("Error executing script:", error)
		return
	}
	println("Executing script:", cmd)

	fmt.Println("Executing script:", path)
	return
}

func parseScripts(path string) {

	jsonFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(jsonFile), &scripts); err != nil {
		panic(err)
	}

	for _, val := range scripts {
		scriptMap[val.ID] = val
	}
}
