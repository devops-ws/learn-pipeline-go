build:
	CGO_ENABLED=0 GOARCH=amd64 go build -o go-server -ldflags "-w"
	chmod u+x go-server

image:
	docker build . -t devopsws/go-server -f Dockerfile.multiStage --build-arg VERSION=dev
run: image
	docker run --rm -p 8080:80 devopsws/go-server
