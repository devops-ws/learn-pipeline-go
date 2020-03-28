FROM alpine:3.11.5

MAINTAINER Rick <rick@jenkins-zh.cn>
LABEL Description="This is a demo for golang HTTP server"

EXPOSE 80

COPY go-server go-server

CMD ["./go-server"]
