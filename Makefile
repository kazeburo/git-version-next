VERSION=0.1.0
LDFLAGS=-ldflags "-w -s -X main.version=${VERSION}"

all: git-version-next

.PHONY: git-version-next

git-version-next: cmd/git-version-next/main.go internal/version/*.go
	go build $(LDFLAGS) -o git-version-next cmd/git-version-next/main.go

linux: cmd/git-version-next/main.go internal/version/*.go
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o git-version-next cmd/git-version-next/main.go

check:
	go test ./...

fmt:
	go fmt ./...

tag:
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin master
