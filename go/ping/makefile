IMAGE := json0/ping
TAG := 1.3

.PHONY: build
build: clean ping-build

.PHONY: clean
clean:
	go clean

ping-build:
	go build ping.go

.built-image: Dockerfile
	touch .built-image

docker-build: .built-image
	docker build -t $(IMAGE) .
	docker tag $(IMAGE) $(IMAGE):$(TAG)

docker-push: .built-image
	docker build -t $(IMAGE) .
	docker tag $(IMAGE) $(IMAGE):$(TAG)
	docker push $(IMAGE):$(TAG)

run:
	docker run $(IMAGE):$(TAG)