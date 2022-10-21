dist/katbox:
	mkdir -p dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o dist/katbox

.PHONY: tailwind
tailwind:
	npm run build

.PHONY: dev
dev: 
	air -c .air.toml

.PHONY: run
run: tailwind
	go run main.go

.PHONY: lint
lint:
	golangci-lint run --fix

clean:
	rm -rf dist/*

.PHONY: image
image: dist/katbox
	docker build -t atechnohazard/katbox .

.PHONY: push
push: image
	docker push atechnohazard/katbox

.PHONY: ent
ent:
	go generate ./ent