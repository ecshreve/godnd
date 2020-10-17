mod-init:
	go mod init github.com/ecshreve/godnd

mod-tidy:
	go mod tidy
	
build:
	go build -o bin/godnd github.com/ecshreve/godnd/cmd/godnd

clean:
	rm pkg/genny/generated_types.go

install:
	go install -i github.com/ecshreve/godnd/cmd/godnd

run-only:
	bin/godnd

run: clean build run-only

test:
	go test github.com/ecshreve/godnd/...

testv:
	go test -v github.com/ecshreve/godnd/...

testc:
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/godnd/...