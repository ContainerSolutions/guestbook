PGIMG := postgres:9.6.1-alpine
OUT := dist/guestbook

build:
	@go build -o $(OUT) .

run: build
	@$(OUT)

pg:
	@docker run -d --name pg -p 5432:5432 -e POSTGRES_PASSWORD=foobar $(PGIMG)

pg-client:
	@docker run --rm -it --entrypoint /usr/local/bin/psql --link pg:pg -e PGPASSWORD=foobar $(PGIMG) -h pg -U postgres
