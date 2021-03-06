default: build

build: dl

dl:
	@godep go build -o dl

run:
	@godep go run main.go

.PHONY: crate
crate:
	@pushd sample; \
	./crate_server.sh; \
	popd
