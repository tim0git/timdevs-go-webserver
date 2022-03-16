##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN go build -o /go-server

##
## Package
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY static static
COPY --from=build /go-server /go-server

EXPOSE 8443

USER nonroot:nonroot

ENTRYPOINT ["/go-server"]