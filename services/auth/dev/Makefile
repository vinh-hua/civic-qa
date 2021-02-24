DOCKER_USER ?= rbayer
IMAGE_NAME = auth
VERSION_TAG = v0.0.0

FULL_IMAGE = $(DOCKER_USER)/$(IMAGE_NAME):$(VERSION_TAG)

docker_cfg:
	@echo "=== CONFIG ==="
	@echo "docker user: $(DOCKER_USER)"
	@echo "set env DOCKER_USER=<YOUR USER> on your system to change"
	@echo "output image: $(FULL_IMAGE)"

build: docker_cfg
	@echo "=== BUILD ==="
	@echo "building and tagging $(IMAGE_NAME):$(VERSION_TAG)"

	cd .. && docker build -t $(FULL_IMAGE) .

	docker push $(FULL_IMAGE)

test:
	@echo "=== TEST ==="
	@echo "Running all tests"
	go test ../...

test_verbose: 
	@echo "=== TEST ==="
	@echo "Running all tests (verbose)"
	go test ../... -v