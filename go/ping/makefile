IMAGE := json0/ping
TAG := 1.1

.PHONY: build
build: clean ping-build

.PHONY: clean
clean:
	go clean

ping-build:
	go build main.go

.built-image: Dockerfile
	docker build -t $(IMAGE) .
	touch .built-image

push: .built-image
	docker tag $(IMAGE) $(IMAGE):$(TAG)
	docker push $(IMAGE):$(TAG)
