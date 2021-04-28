[![Docker Pulls](https://img.shields.io/docker/pulls/devopsws/go-server.svg)](https://hub.docker.com/r/devopsws/go-server/tags)

This is very simple HTTP server written by golang.

# Get started

Start it by following command:

`docker run --rm -p 8080:80 devopsws/go-server:latest`

then you can visit it via: `curl http://localhost:8080`

# kustomize

Please check out [learn-kustomize](https://github.com/devops-ws/learn-kustomize) if you want to this demo into Kubernetes. For example, you can change the docker image tag via [kustomize](https://github.com/kubernetes-sigs/kustomize/), then [ArgoCD](https://github.com/argoproj/argo-cd/) can deploy it to Kubernetes cluster.
