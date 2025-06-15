FROM golang:1.24-alpine AS builder

WORKDIR /usr/local/src

COPY src/go.mod src/go.sum /

RUN go mod download

COPY src ./
RUN go build -o ./bin/app cmd/app/main.go

FROM alpine

COPY --from=builder /usr/local/src/bin/app /usr/local/src/.env /

CMD [ "/app" ]