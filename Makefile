INSTALL_DIR = /usr/local/bin

ifdef prefix
	INSTALL_DIR = $(prefix)
endif

build:
	go build -o bin/check-primer-unicity ./cmd/check-primer-unicity/main.go