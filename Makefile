VERSION=0.0.2
LDFLAGS=-ldflags "-w -s -X main.version=${VERSION}"

all: git-version-next

.PHONY: git-version-next

git-version-next: main.go
	go build $(LDFLAGS) -o git-version-next

linux: main.go
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o git-version-next

check:
	go test ./...

fmt:
	go fmt ./...

tag:
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin master
