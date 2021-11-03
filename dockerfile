FROM golang:1.17.2-alpine3.14

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN go build ./cmd/gochess/gochess.go

EXPOSE 8080

CMD [ "/gochess" ]
