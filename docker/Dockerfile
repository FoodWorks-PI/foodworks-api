# syntax = docker/dockerfile:1-experimental
FROM golang:1.15.1-alpine AS base
WORKDIR /src
ENV GO111MODULE=on

COPY go.* ./
RUN go mod download
RUN go mod verify
COPY . ./

FROM base as build
# TODO: Do we need this?
RUN \
    --mount=type=cache,target=/tmp/.buildx-cache/go-build \
    go build -o /out/server

FROM alpine AS bin-unix
# HACKS
EXPOSE 8080
COPY --from=build /out/server /
CMD ["./server"]
