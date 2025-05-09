RunMonitor:
	go run src/monitor/monitor.go

RunMain:
	go run src/main.go

Run:
	go run src/monitor/monitor.go &
	go run src/main.go 

