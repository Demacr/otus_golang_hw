build:
	go build  -o calendar ./cmd/calendar/...

test:
	go test -race ./... -count=10 -cover

run:
	./calendar --config=configs/config.yaml

lint:
	golangci-lint run ./...

migrate:
	go build -o migrator cmd/migrator/*.go
	./migrator -dir migrations/ "user=calendar dbname=calendar sslmode=disable" up

clean:
	rm -rf *.log
	rm calendar
	rm migrator

.PHONY: build
