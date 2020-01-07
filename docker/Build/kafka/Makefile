#
# Makefile used to build a Kafka Docker image
#
ALPINE_VERSION := 3.10.3
JRE_VERSION    := 11
KAFKA_VERSION  := 2.3.1
SCALA_VERSION  := 2.12
REGISTRY       := my-private-registry.com
REPO           := kafka
NAME           := $(KAFKA_VERSION)-$(SCALA_VERSION)-alpine$(ALPINE_VERSION)

##########

all: build push

build:
	docker build -t $(REGISTRY)/$(REPO):$(NAME) \
	--build-arg ALPINE_VERSION=$(ALPINE_VERSION) \
  	--build-arg JRE_VERSION=$(JRE_VERSION) \
	--build-arg KAFKA_VERSION=$(KAFKA_VERSION) \
	--build-arg SCALA_VERSION=$(SCALA_VERSION)

push:
	docker push $(REGISTRY)/$(REPO):$(NAME)
