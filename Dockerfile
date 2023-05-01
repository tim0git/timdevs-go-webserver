##
## Build
##

FROM golang:1.19.0-buster AS build

ENV GOARCH="arm64"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go get all

COPY *.go ./

RUN go build -o /go-server

##
## Package
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY build build
COPY --from=build /go-server /go-server

EXPOSE 8443

USER nonroot:nonroot

ENTRYPOINT ["/go-server"]
