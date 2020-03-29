build:
	GOOS=linux GOARCH=amd64 go build -o go-server
	chmod u+x go-server

image:
	docker build . -t devopsws/go-server
