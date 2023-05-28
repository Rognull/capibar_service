FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY api ./api
COPY cmd ./cmd
COPY internals ./internals

RUN CGO_ENABLED=0 GOOS=linux go build -v ./cmd/main.go

CMD ["./main"]