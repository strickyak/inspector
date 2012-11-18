all:
	GOPATH=$$PWD  go build -v -x src/main.go 

run:
	GOPATH=$$PWD  go run src/main.go 9999
