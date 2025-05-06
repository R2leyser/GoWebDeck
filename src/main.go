package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
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
	nextID    = 1
	postsMu   sync.Mutex
)

func main() {
	http.HandleFunc("/scripts/", scriptHandler)
	http.HandleFunc("/", frontendHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Script added:", scriptMap[nextID-1].Path)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	parseScripts(os.Getenv("HOME") + "/.config/gowebdeck/scripts.json")

    w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(http.StatusOK)

    fmt.Fprintln(w, "<!DOCTYPE html>")
    fmt.Fprintln(w, "<head>")
    fmt.Fprintln(w, "<link rel=\"stylesheet\" type=\"text/css\" href=\"/static/style.css\">")
    fmt.Fprintln(w, "</head>")
    fmt.Fprintln(w, "<html><body>")
    fmt.Fprintln(w, "<h1>Script Runner</h1>")
    fmt.Fprintln(w, "<ul>")

    for _, script := range scriptMap {
        fmt.Fprintf(w, "<div class=\"script-button\" data-id=\"%d\">", script.ID)
        fmt.Fprintf(w, "<p>%s</p>", script.Description)
        fmt.Fprintf(w, "</div>")
    }

    fmt.Fprintln(w, "</body></html>")
    fmt.Fprintln(w, "<script src=\"/static/index.js\"></script>")
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
	var scripts []Script

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
