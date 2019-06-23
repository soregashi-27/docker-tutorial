FROM quay.io/cybozu/golang:1.12-bionic AS build

COPY main.go /work/main.go
WORKDIR /work
RUN CGO_ENABLED=0 go build -o server ./main.go

FROM scratch
COPY --from=build /work/server /server
EXPOSE 8000
ENTRYPOINT ["/server"]
