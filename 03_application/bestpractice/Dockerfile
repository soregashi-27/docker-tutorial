FROM quay.io/cybozu/golang:1.12-bionic AS build

COPY src /work/src
WORKDIR /work/src
RUN go mod download
RUN CGO_ENABLED=0 go install ./...

FROM scratch
COPY --from=build /go/bin /

ENV PORT=8080
EXPOSE $PORT

HEALTHCHECK --interval=3s --timeout=5s CMD [ "/healthcheck" ]
ENTRYPOINT ["/server"]
