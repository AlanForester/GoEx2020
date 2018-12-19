# ./Makefile
VERSION := $(shell date +'%Y%m%d%H').$(shell git rev-parse --short=8 HEAD)
.PHONY: proto data build

NAME := $(shell echo api)

GOPWD := $(shell pwd)
GOBASEDIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../..)
#GOHOMEDIR := $(abspath $(HOME)/go)

GOPATH = $(GOPWD):$(GOBASEDIR)
GOBIN = $(GOHOMEDIR)/bin
PROTOC := `which protoc`
$(info root makefile GOPATH=$(GOPATH))
$(info root makefile GOBIN=$(GOBIN))

gen: ./proto/*.proto
	@echo "Compiling proto file $(^)"
#Swagger
	$(shell $(PROTOC) -I/usr/local/opt/protobuf/include -I. -I$(GOPATH)/src \
	-I$(GOBASEDIR)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:. $(^) )

#GRPC PB
	$(shell $(PROTOC) -I/usr/local/opt/protobuf/include -I. -I$(GOPATH)/src \
	-I$(GOBASEDIR)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:. $(^) )

#Micro
	$(shell $(PROTOC) -I/usr/local/opt/protobuf/include -I. -I$(GOPATH)/src \
	--micro_out=. --go_out=plugins=micro:. $(^) )

#Gateway
	$(shell $(PROTOC) -I/usr/local/opt/protobuf/include -I. -I$(GOPATH)/src \
      -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --grpc-gateway_out=logtostderr=true:. $(^) )

	@echo "Generating is done."
	exit 0;

proto:
	for d in gateway srv; do \
		for f in $$d/proto/**/*.proto; do \
			echo compilling: $$f; \
			${PROTOC} -I/usr/local/opt/protobuf/include --swagger_out=logtostderr=true:. $$f; \
			${PROTOC} -I/usr/local/opt/protobuf/include --go_out=plugins=grpc:. $$f; \
			${PROTOC} -I/usr/local/opt/protobuf/include --micro_out=. --go_out=:. $$f; \
			${PROTOC} -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I. --grpc-gateway_out=logtostderr=true:. $$f; \
			echo compiled: $$f; \
		done \
	done


docker:
	docker-compose build
	docker-compose up

all:
	@echo "Project:" $(NAME) $(VERSION)

builds: $(NAME)
$(NAME): *.go
	go build -o ./deploy/build/$(NAME) -v

run:
	go run main.go
