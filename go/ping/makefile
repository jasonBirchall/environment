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

docker-push: .built-image
	docker build -t $(IMAGE) .
	docker tag $(IMAGE) $(IMAGE):$(TAG)
	docker push $(IMAGE):$(TAG)

shell:
	docker build -t $(IMAGE):$(TAG) .
	docker run -it --rm $(IMAGE):$(TAG) sh

run:
	docker build -t $(IMAGE):$(TAG) .
	docker run $(IMAGE):$(TAG)