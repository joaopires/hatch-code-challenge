test:
	go test -v ./...

build:
	go build -o hatch ./cmd/hatch

install:
	go install ./cmd/hatch