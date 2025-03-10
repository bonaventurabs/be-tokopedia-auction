# Builder
FROM golang:1.21.5-alpine3.17 as builder

RUN apk update && apk upgrade && \
  apk --update add git make bash build-base

WORKDIR /app

COPY . .

COPY ./config.ini .

RUN make build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
  apk --update --no-cache add tzdata && \
  mkdir /app 

WORKDIR /app 

EXPOSE 8080

COPY --from=builder /app/engine /app/
COPY --from=builder /app/config.ini /app/

CMD /app/engine