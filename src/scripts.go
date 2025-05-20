package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func scriptsInit() {
	http.HandleFunc("/scripts/", scriptHandler)
}

func scriptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handlePostScript(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePostScript(w http.ResponseWriter, r *http.Request) {
    // Parse the script ID from the URL, e.g., /scripts/1/on
    on := strings.Contains(r.URL.Path, "/on")
    off := strings.Contains(r.URL.Path, "/off")

    idStr := r.URL.Path[len("/scripts/"):]
    if strings.Contains(idStr, "/on") {
        idStr = strings.TrimSuffix(idStr, "/on")
    } else if strings.Contains(idStr, "/off") {
        idStr = strings.TrimSuffix(idStr, "/off")
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid script ID", http.StatusBadRequest)
        return
    }

	p, ok := scriptMap[id]
	if !ok {
		http.Error(w, "Script not found", http.StatusNotFound)
		return
	}

	fmt.Println(on)
	fmt.Println(off)


    // fix this
	if on {
		go executeScript(p.Path + "/on")
	} else if off {
		go executeScript(p.Path + "/off")
	} else {
		go executeScript(p.Path)
	}
	http.Error(w, "Script execution started", http.StatusAccepted)
}

func executeScript(path string) {
    exec.Command("notify-send", "Executing script", path).Run()

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
