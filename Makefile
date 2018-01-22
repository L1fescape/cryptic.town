PACKAGE = cryptic.town

GO = go
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: clean build run

.PHONY: clean
clean: ; $(info $(M) cleaning)	@
	@rm -rf rpc/out
	@rm -rf web/dist

.PHONY: build
build: build-proto build-web

.PHONY: build-proto
build-proto: ; $(info $(M) generating proto files)	@
	@mkdir -p ./rpc/out
	@protoc -I ./rpc --twirp_out=./rpc/out --go_out=./rpc/out ./rpc/home.proto

.PHONY: build-web
build-web: clean; $(info $(M) building frontend)	@
	@cd web && yarn build

.PHONY: run
run: ; $(info $(M) running service)	@
	$(GO) run main.go

