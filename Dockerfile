FROM golang:1.17-alpine as builder

RUN apk add build-base

WORKDIR /opt
COPY . .
ENV PATH="$PATH:/opt"

RUN go build ./...
RUN go test ./...

FROM golang:1.17-alpine
WORKDIR /opt

COPY --from=builder /opt/birdpedia /opt
COPY --from=builder /opt/birdpedia.sh /opt
COPY --from=builder /opt/assets /opt/assets

ENTRYPOINT ["/opt/birdpedia.sh"]