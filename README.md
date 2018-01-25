# cryptic.town

## Install

```
$ go get
$ npm i -g yarn
```

## Run

Build and run everything

```
$ SLACK_TOKEN=[token] make
```

Skip build and run

```
$ SLACK_TOKEN=[token] make run
```

## Develop

Install [protobuf](https://github.com/google/protobuf)

```
$ go get github.com/twitchtv/twirp/protoc-gen-twirp
$ go get github.com/golang/protobuf/protoc-gen-go
$ make dev
```
