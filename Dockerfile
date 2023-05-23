FROM golang:alpine AS capy
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -o /main cmd/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /build
COPY --from=capy /main /main
EXPOSE 8080

USER noroot:noroot

ENTRYPOINT ["/main"]