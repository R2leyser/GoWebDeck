package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
)

type Script struct {
    ID   int    `json:"id"`
    Path string `json:"path"`
}

var (
    scripts   = make(map[int]Script)
    nextID  = 1
    postsMu sync.Mutex
)

func main() {
    http.HandleFunc("/scripts/", scriptHandler)

    scripts[nextID] = Script{ID: nextID, Path: "/path/to/script1"}
    nextID++

    fmt.Println("Script added:", scripts[nextID-1].Path)

    fmt.Println("Server is running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))

    // Example of adding a script 
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

    p, ok := scripts[id]
    if !ok {
        http.Error(w, "Script not found", http.StatusNotFound)
        return
    }

    fmt.Println("Script ID:", p.ID)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode("Something")
}

func handleGetScript(w http.ResponseWriter, r *http.Request, id int) {
    postsMu.Lock()
    defer postsMu.Unlock()

    // If you use a two-value assignment for accessing a
    // value on a map, you get the value first then an
    // "exists" variable.
    _, ok := scripts[id]
    if !ok {
        http.Error(w, "Script not found", http.StatusNotFound)
        return
    }

    delete(scripts, id)
    w.WriteHeader(http.StatusOK)
}
