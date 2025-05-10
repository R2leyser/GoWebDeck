runmonitor:
	go run src/monitor/*.go

runmain:
	go run src/*.go

run:
	go run src/monitor/*.go &
	go run src/*.go 

