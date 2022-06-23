FROM golang

ENV GO111MODULE=on
RUN mkdir /build
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

CMD go test -v ./...

