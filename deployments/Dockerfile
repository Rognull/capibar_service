FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY api ./api
COPY cmd ./cmd
COPY internals ./internals
COPY ./Makefile ./

RUN CGO_ENABLED=0 GOOS=linux make build

CMD ["./main"]