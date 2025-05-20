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
	var on = false
	var off = false

	fmt.Println(r.URL.Path)
	if strings.Contains(r.URL.Path, "on"){
		on = true
	} else if strings.Contains(r.URL.Path, "off") {
		off = true
	}

	fmt.Println(r.URL.Path[len("/scripts/"):])
	id, err := strconv.Atoi(r.URL.Path[len("/scripts/"):])
	if err != nil {
		http.Error(w, "Invalid Script ID", http.StatusBadRequest)
		return
	}

	fmt.Println(id)
	if r.Method == "POST" {
		handlePostScript(w, r, id, on, off)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePostScript(w http.ResponseWriter, r *http.Request, id int, on bool, off bool) {
	fmt.Println(r.URL.Path)
	p, ok := scriptMap[id]
	if !ok {
		http.Error(w, "Script not found", http.StatusNotFound)
		return
	}

	fmt.Println(on)
	fmt.Println(off)

	if on {
		fmt.Println("post on")
	} else if off {
		fmt.Println("post off")
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
