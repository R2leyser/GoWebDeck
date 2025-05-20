	package main

// this is only needed because my server isn't portforwarded or on my network and I don't trust myself
import (
	"fmt"
	"io"
	"net/http"
)

func servidorRamHandler(w http.ResponseWriter, r *http.Request) {
	serverIP := "http://172.26.67.173:8000"
	serverURL := fmt.Sprintf("%s/ram", serverIP)
	makeCrossOrigin(&w)

	res, err := http.Get(serverURL)
	if err != nil {
		fmt.Println("Couldn't connect to server")
	} else {
		io.Copy(w, res.Body)
	}

}

func servidorCpuHandler(w http.ResponseWriter, r *http.Request) {
	serverIP := "http://172.26.67.173:8000"
	serverURL := fmt.Sprintf("%s/cpu", serverIP)
	makeCrossOrigin(&w)

	res, err := http.Get(serverURL)
	if err != nil {
		fmt.Println("Couldn't connect to server")
	} else {
		io.Copy(w, res.Body)
	}
}
