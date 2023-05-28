FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

#COPY api ./
#COPY cmd ./
#COPY internals ./
COPY . .

RUN go build cmd/main.go

CMD ["/main"]