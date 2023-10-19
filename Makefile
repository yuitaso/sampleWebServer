REPOSITORY_ROOT = $(shell git rev-parse --show-toplevel)

.PHONY: init clear/db sqlite

init: sqlite
	go run dev/init.go

server:
	go run src/main.go

clear/db:
	rm -rf $(REPOSITORY_ROOT)/sqlite/

sqlite:
	mkdir $(REPOSITORY_ROOT)/sqlite
