FROM golang:1.20.4-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
ENV SERV_PORT=8080
ENV SERV_DBUSER=postgres
ENV SERV_DBPASS=tmppass
ENV SERV_DBHOST=195.133.197.62
ENV SERV_DBPORT=3030
ENV SERV_DBNAME=capybaras
#COPY api ./
#COPY cmd ./
#COPY internals ./
COPY . .

RUN go build cmd/main.go

CMD ["./main"]