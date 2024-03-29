[![Docker Pulls](https://img.shields.io/docker/pulls/devopsws/go-server.svg)](https://hub.docker.com/r/devopsws/go-server/tags)

This is very simple HTTP server written by golang.

## Get started
Start it by following command:

`docker run --rm -p 8080:80 devopsws/go-server:latest`

then you can visit it via: `curl http://localhost:8080`

## APIs
This project offers the following APIs:

| Path | Description |
|---|---|
| `/` | Print a hello world message. |
| `/version` | Print the version of this project. |

## Flux
Please go into the directory [flux](flux) if you like to try it.

## Argo Rollouts
See also [Argo Rollouts](https://github.com/devops-ws/argo-rollouts-guide) related files [here](argo/rollouts/).
