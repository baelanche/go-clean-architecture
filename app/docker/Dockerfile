FROM golang:1.22-alpine

WORKDIR /go

COPY . .

RUN go mod download

RUN go build -o go-clean-architecture ./app/cmd/main.go

EXPOSE 8080

ENTRYPOINT [ "./go-clean-architecture" ]