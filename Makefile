GO := $(shell which go)
GLIDE := $(shell which glide)
PROTOC := $(shell which protoc)

all:
	$(GLIDE) install

	$(PROTOC) -I. -I/usr/local/include -I $(GOPATH)/src \
	--go_out=plugins=grpc:. \
	--govalidators_out=. \
	proto/greet.proto

	$(GO) build --buildmode=plugin -o plugin.so plugin/main.go

	$(GO) run main.go
