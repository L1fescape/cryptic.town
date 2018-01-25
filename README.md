# cryptic.town

## Install

Install [protobuf](https://github.com/google/protobuf)

```
$ go get
$ go get github.com/twitchtv/twirp/protoc-gen-twirp
$ go get github.com/golang/protobuf/protoc-gen-go
$ npm i -g yarn
```

## Run


Build everything and run

```
$ SLACK_TOKEN=[token] make
```

Skip build and run

```
$ SLACK_TOKEN=[token] make run
```

## Develop

```
$ make dev
```
