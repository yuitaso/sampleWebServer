REPOSITORY_ROOT = $(shell git rev-parse --show-toplevel)
DB_PATH := $(REPOSITORY_ROOT)/sqlite/api.db

.PHONY: check install init clear/db

check:
	@go version | grep '1.21.3' || echo 'go1.21.3 reqired.'   

install: 
	go mod tidy

init: sqlite
	go run dev/init.go

server:
	go run src/main.go

clean:
	rm -rf $(REPOSITORY_ROOT)/sqlite/*

sqlite:
	mkdir $(REPOSITORY_ROOT)/sqlite

reset: clean init

# test
.PHONY: test test/server test/init test/clean test/reset

test:
	go test test/integrationTest/integration_test.go

test/server: test/reset
	go run src/main.go -e test

test/init: sqlite
	go run dev/init.go -e test

test/seed:
	sqlite3 ./sqlite/test.db < test/integrationTest/seeds.sql

test/clean:
	rm -f sqlite/test.db

test/reset: test/clean test/init test/seed
