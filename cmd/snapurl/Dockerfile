#
# Build
#
FROM golang:1.13-alpine AS build

# Maintainer Info
LABEL maintainer="Bruno Luiz Silva <contact@brunoluiz.net>"

## Install build utilities
ENV OS_PACKAGES  "git make ca-certificates"
RUN apk add ${OS_PACKAGES} --update --no-cache

## Private repo if required
ARG GITHUB_TOKEN
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

WORKDIR /app

## Copy go mod and sum files
COPY go.mod go.sum ./

## Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

## Copy the source from the current directory to the Working Directory inside the container
COPY . .
RUN make build-cli build-grpc

#
# Runtime
# Based on: https://github.com/chromedp/docker-headless-shell
#
FROM chromedp/headless-shell:78.0.3904.34

ENV PATH=$PATH:/headless-shell

RUN apt-get update
RUN apt-get install -y tini ca-certificates
RUN update-ca-certificates 2>/dev/null || true

WORKDIR /
COPY --from=build /app/bin/snapurl ./snapurl

ENTRYPOINT ["tini", "--"]
CMD ["/snapurl"]

