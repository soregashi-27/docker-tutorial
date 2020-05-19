# Stage1: build from source
FROM quay.io/cybozu/golang:1.13-bionic AS build

COPY main.go /work/main.go

WORKDIR /work
RUN CGO_ENABLED=0 go build -o server ./main.go

# Stage2: setup runtime container
FROM scratch

COPY --from=build /work/server /server
EXPOSE 8000

ENTRYPOINT ["/server"]
