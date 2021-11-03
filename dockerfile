FROM golang:1.17.2-alpine3.14

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build .\cmd\backend\main.go -o /gochess

EXPOSE 8080

CMD [ "/gochess" ]
