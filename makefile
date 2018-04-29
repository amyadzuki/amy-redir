.POSIX:
GO=go
SHELL=/bin/sh
amyredir: ; $(GO) build -o $@ src/*.go
