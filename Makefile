dist/katbox:
	mkdir -p dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o dist/katbox

.PHONY: dev
dev:
	air -c .air.toml

.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	golangci-lint run --fix

clean:
	rm -rf katbox

.PHONY: image
image: katbox
	docker build -t atechnohazard/katbox .