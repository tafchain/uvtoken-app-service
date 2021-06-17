FROM golang:alpine as builder

ENV GO111MODULE=on
COPY .  /gin-wallet-service
RUN cd /gin-wallet-service/server && go build -o server .

FROM alpine:latest
#RUN apk add --no-cache ca-certificates
ENV GIN_MODE=release
WORKDIR /app
COPY --from=builder  /gin-wallet-service/server/server     .
COPY --from=builder  /gin-wallet-service/server/config.yaml .

WORKDIR /app
ENTRYPOINT ["/app/server"]