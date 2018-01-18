PACKAGE = cryptic.town

GO = go
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: clean genrpc run

.PHONY: clean
clean: ; $(info $(M) cleaning)	@
	@rm -rf rpc/out

.PHONY: genrpc
genrpc: clean; $(info $(M) generating proto files)	@
	$Q @mkdir ./rpc/out
	$Q @protoc -I ./rpc --twirp_out=./rpc/out --go_out=./rpc/out ./rpc/home.proto

.PHONY: run
run: ; $(info $(M) running service)	@
	$Q $(GO) run main.go

