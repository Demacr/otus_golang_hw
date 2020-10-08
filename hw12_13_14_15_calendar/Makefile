build:
	go build  -o calendar ./cmd/calendar/...

test:
	go test -race ./... -count=10 -cover

run:
	./calendar --config=configs/config.yaml

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.30.0

lint: install-lint-deps
	golangci-lint run ./...

migrate:
	go build -o migrator cmd/migrator/*.go
	./migrator -dir migrations/ "user=calendar dbname=calendar sslmode=disable" up

clean:
	rm -rf *.log
	rm calendar
	rm migrator

.PHONY: build
