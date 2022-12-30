# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /slack-gpt-bot

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app/.env /.env
COPY --from=build /slack-gpt-bot /slack-gpt-bot


ENTRYPOINT ["/slack-gpt-bot"]