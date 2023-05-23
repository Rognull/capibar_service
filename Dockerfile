FROM golang:alpine AS capy
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -o /main cmd/main.go

FROM gcr.io/distroless/base-debian10

RUN apt-get update
RUN apt-get -y install bash

WORKDIR /build
COPY --from=capy /main /main
EXPOSE 8080
ENV SERV_PORT=8080
ENV SERV_DBUSER=postgres
ENV SERV_DBPASS=tmppass
ENV SERV_DBHOST=195.133.197.62
ENV SERV_DBPORT=3030
ENV SERV_DBNAME=capybaras
USER noroot:noroot

ENTRYPOINT ["/main"]