FILENAME := ''

fetch_migrate:
ifeq ($(shell uname), Darwin)
	curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.darwin-amd64.tar.gz | tar xvz

	mv migrate db
else ifeq ($(shell uname), Linux)
	curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz

	mv migrate db
else ifeq ($(OS), Windows_NT)
	curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.windows-amd64.zip -o db/migrate.zip
	cd db && unzip migrate.zip
endif

# run migration down and up
run_migration:
	db/migrate -path=db/migrations/ -database "postgres://postgres:postgres123@localhost:5432/books?sslmode=disable" down
	db/migrate -path=db/migrations/ -database "postgres://postgres:postgres123@localhost:5432/books?sslmode=disable" up

create_migrate:
ifeq ($(FILENAME), '')
	$(error FILENAME cannot be empty)
endif
ifeq ($(shell uname), Darwin || Linux)
	./db/migrate create -ext sql -dir db/migrations $(FILENAME)
else ifeq ($(OS), Windows_NT)
	./db/migrate.exe create -ext sql -dir db/migrations FILENAME
endif

