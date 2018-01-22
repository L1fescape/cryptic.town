PACKAGE = cryptic.town

GO = go
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: clean genrpc run

.PHONY: clean
clean: ; $(info $(M) cleaning)	@
	@rm -rf rpc/out

.PHONY: genrpc
genrpc: clean; $(info $(M) generating proto files)	@
	@mkdir ./rpc/out
	@protoc -I ./rpc --twirp_out=./rpc/out --go_out=./rpc/out ./rpc/home.proto

.PHONY: run
run: ; $(info $(M) running service)	@
	$(GO) run main.go

