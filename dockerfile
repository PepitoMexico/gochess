FROM golang:1.17.2-alpine3.14

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/backend/*.go ./
COPY cmd/frontend/*.go ./

RUN go build -o /gochess

EXPOSE 8080

CMD [ "/gochess" ]
