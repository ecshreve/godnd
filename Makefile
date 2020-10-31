mod-init:
	go mod init github.com/ecshreve/godnd

mod-tidy:
	go mod tidy
	
build:
	go build -o bin/godnd github.com/ecshreve/godnd/cmd/godnd
	go build -o bin/genny github.com/ecshreve/godnd/cmd/genny

install:
	go install -i github.com/ecshreve/godnd/cmd/godnd
	go install -i github.com/ecshreve/godnd/cmd/genny

run-only-godnd:
	bin/godnd

run-only-genny:
	bin/genny

run-godnd: build run-only-godnd

run-genny: build run-only-genny

test:
	go test github.com/ecshreve/godnd/...

testv:
	go test -v github.com/ecshreve/godnd/...

testc:
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/godnd/...