build:
	mkdir -p bin
	go build -o bin/monitor src/monitor/*.go
	go build -o bin/main src/*.go

run:
	go run src/monitor/*.go &
	go run src/*.go 

runmonitor:
	go run src/monitor/*.go

runmain:
	go run src/*.go



