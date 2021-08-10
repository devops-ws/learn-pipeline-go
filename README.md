[![Docker Pulls](https://img.shields.io/docker/pulls/devopsws/go-server.svg)](https://hub.docker.com/r/devopsws/go-server/tags)

This is very simple HTTP server written by golang.

# Get started
Start it by following command:

`docker run --rm -p 8080:80 devopsws/go-server:latest`

then you can visit it via: `curl http://localhost:8080`

# Flux
Please go into the directory [flux](flux) if you like to try it.

# Podman

You can build the image via [podman](https://github.com/containers/podman):

`podman build . --events-backend=file -f Dockerfile.multiStage -t surenpi/learn-pipeline-go`
