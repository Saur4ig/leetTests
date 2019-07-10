.PHONY: lint, deps, test, fmtall, all, outdated, tags, push

deps:
# install linter
	if [ -z `which golangci-lint` ]; then go get -u github.com/golangci/golangci-lint/cmd/golangci-lint; fi
	if [ -z `which go-mod-outdated` ]; then go get -u github.com/psampaz/go-mod-outdated; fi

lint:
	golangci-lint run

test:
	go test -count=1 ./... -cover

fmtall:
	go fmt ./...

all:
	make fmtall
	make lint
	make test

outdated:
	go list -u -m -json all | go-mod-outdated -update

tags:
	git push --tags

push:
	git push --follow-tags