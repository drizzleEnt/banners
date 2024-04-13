FROM golang:1.21.1-alpine AS builder

COPY . /github.com/drizzleent/banners/source/
WORKDIR /github.com/drizzleent/banners/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/drizzleent/banners/source/.env .
COPY --from=builder /github.com/drizzleent/banners/source/bin/crud_server .

EXPOSE 8080
CMD [ "./crud_server" ]