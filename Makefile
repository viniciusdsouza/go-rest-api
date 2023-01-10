TARGET_FILE:=${shell head -n1 go.mod | sed -r 's/.*\/(.*)/\1/g' }
BUILD_DIR=.build
COVER_PROFILE_FILE="${BUILD_DIR}/go-cover.tmp"
.PHONY: clean mk-build-dir update-deps build-deps build build-webserver docker-build clean-test test cover-html docs

clean:
	rm -rf $(TARGET_FILE) $(BUILD_DIR)

############## build tasks

mk-build-dir:
	@mkdir -p ${BUILD_DIR}

update-deps:
	@GOPRIVATE=github.com/dock-tech go get -u -d -v ./...
	go mod tidy
	
build-deps:
	@GOPRIVATE=github.com/dock-tech go get -d -v ./...
	go mod tidy

build: build-deps build-webserver

build-webserver:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o ./${BUILD_DIR}/webserver ./cmd/webserver

docker-build: 
	docker build . -t $(TARGET_FILE)

############## test tasks

clean-test:
	@go fmt ./...
	@go clean -testcache

test: clean-test
	go test ./... --cover

cover-html: mk-build-dir clean-test
	go test -p 1 -coverprofile=${COVER_PROFILE_FILE} ./... ; echo
	go tool cover -html=${COVER_PROFILE_FILE}

############## docs
docs:
	swag init --dir cmd/webserver --parseDependency --parseInternal --parseDepth 1
	rm docs/docs.go