REPOSITORY_ROOT = $(shell git rev-parse --show-toplevel)


init: sqlite
	go run dev/init.go

clear/db:
	rm -rf $(REPOSITORY_ROOT)/sqlite/

sqlite:
	mkdir $(REPOSITORY_ROOT)/sqlite
