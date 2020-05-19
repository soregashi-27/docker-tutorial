# Stage1: build from source
FROM quay.io/cybozu/golang:1.13-bionic AS build

COPY src /work/src
WORKDIR /work/src

RUN go mod download
RUN CGO_ENABLED=0 go install ./...

# Stage2: setup runtime container
FROM scratch

COPY --from=build /go/bin /

ENV PORT=9000
EXPOSE $PORT

ENTRYPOINT ["/server"]
