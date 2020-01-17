FROM golang:1.12 as builder
ENV GO111MODULE=on
WORKDIR /root
# cache pkg
COPY go.mod .
COPY go.sum .
RUN go mod download

# build
COPY cmd ./cmd
COPY pkg ./pkg
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install ./cmd/...

FROM alpine:3.10.0

COPY --from=builder /go/bin /usr/local/bin
# add static
COPY templates ./templates

