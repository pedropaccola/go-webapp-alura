build:
	@go build -o bin/webapp

# db:
# 	@docker run --name some-postgres -e POSTGRES_PASSWORD=webapp -p 5432:5432 -d postgres

run: build
	@./bin/webapp

test:
	@go test -v ./...