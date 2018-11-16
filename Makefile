PKG := github.com/w32blaster/bot-scaffolding
PKG_LIST=$(shell go list ${PKG}/... | grep -v /vendor/)
FILES := $$(find . -name '*.go' | grep -vE 'vendor')

test:
	@go test -v -race ${PKG_LIST}

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags "-s -w" -o bot-scaffolding  github.com/w32blaster/bot-scaffolding

generateInit:
	@go run main.go -o init -t ./sandbox
	@cd sandbox
	@go build -v .

regenerateImplementation:
	@go run main.go -o generate -t ./sandbox

.POHNY: test build generateTest