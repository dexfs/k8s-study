VERSION = 6
docker-build:
	@docker build -t dexfs/hello-go:v${VERSION} .
	@docker build -t dexfs/hello-go:latest .

docker-push:
	@docker push dexfs/hello-go:v${VERSION}
	@docker push dexfs/hello-go:latest

run:
	@go run server.go