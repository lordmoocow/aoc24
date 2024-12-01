ARG GO_VERSION="1.23"
ARG ALPINE_VERSION="3.20"
ARG GO_RELEASER_VERSION="2.4.8"
ARG REPO_URL="github.com/lordmoocow/aoc24"

FROM docker.io/golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS build-base
ARG GO_RELEASER_VERSION
ADD https://github.com/goreleaser/goreleaser/releases/download/v${GO_RELEASER_VERSION}/goreleaser_${GO_RELEASER_VERSION}_x86_64.apk /goreleaser.apk
RUN apk add --allow-untrusted /goreleaser.apk

FROM build-base AS build
ARG REPO_URL
WORKDIR /go/src/${REPO_URL}
RUN --mount=rw,target=. \ 
    --mount=type=cache,target=/root/.cache/go-build \
    goreleaser build --single-target --auto-snapshot -o /usr/bin/aoc

FROM docker.io/golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS run
COPY --from=build /usr/bin/aoc /usr/bin/aoc
ENTRYPOINT ["aoc"]
