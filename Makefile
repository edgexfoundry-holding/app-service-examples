.PHONY: build clean

GO=CGO_ENABLED=1 GO111MODULE=on go

APP_SERVICES=app-services/*

GIT_SHA=$(shell git rev-parse HEAD)

.PHONY: build $(APP_SERVICES)

build: $(APP_SERVICES)

$(APP_SERVICES):
	$(GO) build $(GOFLAGS) -o $@/app-service ./$@

clean:
	rm -f app-services/*/app-service

docker:
	docker build \
	    --build-arg http_proxy \
	    --build-arg https_proxy \
		-f app-services/simple-filter-xml/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/docker-simple-filter-xml:$(GIT_SHA) \
		-t edgexfoundry/docker-simple-filter-xml:dev \
		.
