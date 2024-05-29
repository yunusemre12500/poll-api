ARG GO_VERSION="1.22.3"

FROM golang:${GO_VERSION}-alpine3.19 AS base

WORKDIR /usr/src

COPY . .

FROM base AS build

RUN --mount=type=cache,target=$HOME/go/pkg/mod/ \
	go build -mod=vendor -o=./poll-api /usr/src/cmd/poll-api/main.go

FROM gcr.io/distroless/static-debian12

COPY --chown=nonroot:nonroot --from=build \
	/usr/src/poll-api /usr/local/bin/poll-api

USER nonroot:nonroot

CMD [ "/usr/local/bin/poll-api" ]