package main

import (
	"encoding/json"
	"fmt"
    "log"
	"net/http"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/mem"
)

func main () {

    http.HandleFunc("/ram/", ramMonitorHandler)
    http.HandleFunc("/cpu/", cpuMonitorHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func cpuMonitorHandler(w http.ResponseWriter, r *http.Request) {
    percentages, err := cpu.Percent(0, false)
    if err != nil {
        log.Fatal(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(fmt.Sprintf(`{percent:%.2f}`, percentages[0]));
}

func ramMonitorHandler(w http.ResponseWriter, r *http.Request) {
    virtualMem, err := mem.VirtualMemory()
    if err != nil {
        log.Fatal(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(fmt.Sprintf(`{percent:%.2f, mbUsado:%f, mbTotal:%f}`, 
        virtualMem.UsedPercent, 
        float64(virtualMem.Used)/float64(1000000), 
        float64(virtualMem.Total)/float64(1000000)));
}


