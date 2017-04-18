PGIMG    := postgres:9.6.1-alpine
OUT      := dist/guestbook
OUTOSX   := $(OUT).darwin64
OUTLINUX := $(OUT).linux-amd64

.PHONY: dist
dist: build-osx build-linux

build-osx:
	@go build -o $(OUTOSX) .

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o $(OUTLINUX) .

run:
	@go run main.go

pg:
	@docker run -d --name pg -p 5432:5432 -e POSTGRES_PASSWORD=foobar $(PGIMG)

pg-client:
	@docker run --rm -it --entrypoint /usr/local/bin/psql --link pg:pg -e PGPASSWORD=foobar $(PGIMG) -h pg -U postgres
