FROM golang 

LABEL maintainer="dgaye-pka-mnom"

COPY . /go/src/app

WORKDIR /go/src/app

EXPOSE 8080

CMD go run ./cmd/forum-app/main.go