FROM golang:1.17.0-alpine

WORKDIR /app

COPY . /app

RUN go mod tidy

CMD [ "go", "run", "main.go" ]