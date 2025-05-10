package main

// this is only needed because my server isn't portforwarded or on my network and I don't trust myself
import (
	// "fmt"
	"net/http"
)

var (
	// serverIP = "http://127.0.0.1:8000"
)

func servidorRamHandler ( w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.WriteHeader(http.StatusOK)
	

	// serverURL := fmt.Sprintf("{}/ram", serverIP)
	// res, err := http.Get(serverURL)

	// if err != nil {
		
	// }


}


func servidorCpuHandler ( w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.WriteHeader(http.StatusOK)

	// serverURL := fmt.Sprintf("{}/cpu", serverIP)

	// res, err := http.Get(serverURL)
	// if err != nil {
		
	// }
	
}
