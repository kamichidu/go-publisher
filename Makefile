GO=go
LDFLAGS=-ldflags "-X main.appVersion=$$(git describe --tags)"
PKGS=$$(glide novendor)

.PHONY: build
build: generate
	${GO} build ${TAGS} ${LDFLAGS} .

.PHONY: install
install: generate
	${GO} install ${TAGS} ${LDFLAGS} .

.PHONY: generate
generate:
	${GO} generate ${TAGS} ${PKGS}

.PHONY: test
test: generate
	${GO} test ${PKGS}

.PHONY: deps
deps:
	go get -v github.com/jteeuwen/go-bindata/...
	go get -v github.com/Masterminds/glide
	glide install
