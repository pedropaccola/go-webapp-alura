build:
	@go build -o bin/webapp

dbrun:
	@docker run --name some-postgres -e POSTGRES_PASSWORD=webapp -p 5432:5432 -d postgres
	
dbstop:
	@docker stop some-postgres
	@docker rm some-postgres

run: build
	@./bin/webapp

test:
	@go test -v ./...
