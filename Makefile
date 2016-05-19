default: build

build: dl

dl:
	@go build -o dl

run:
	@go run main.go
