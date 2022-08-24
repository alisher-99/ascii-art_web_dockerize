FROM golang:1.17
LABEL maintainer="alisher-99"

ENV APPNAME="ascii-server"

WORKDIR /ascii-art-docker
COPY . . 

RUN go build -o APPNAME cmd/app/main.go

EXPOSE 8181
ENTRYPOINT [ "./APPNAME" ]
