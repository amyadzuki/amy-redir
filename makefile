.POSIX:

GO=go
PORT=8080
SHELL=/bin/sh

all: amyredir .phony

amyredir: ; $(GO) build -o $@ src/*.go

run: amyredir ; ./amyredir -p $(PORT)

.phony:
