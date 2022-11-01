FROM golang:1.18-bullseye

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o shorty ./cmd/app/main.go

CMD ["./shorty"]