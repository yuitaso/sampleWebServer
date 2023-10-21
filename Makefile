REPOSITORY_ROOT = $(shell git rev-parse --show-toplevel)
DB_PATH := $(REPOSITORY_ROOT)/sqlite/api.db

.PHONY: check install init clear/db sqlite

check:
	@go version | grep '1.21.3' || echo 'go1.21.3 reqired.'   

install: 
	go mod tidy

init: sqlite
	go run dev/init.go -dbPath='$(DB_PATH)'

server:
	go run src/main.go -dbPath=$(DB_PATH)

clear:
	rm -rf $(REPOSITORY_ROOT)/sqlite/

sqlite:
	mkdir $(REPOSITORY_ROOT)/sqlite
