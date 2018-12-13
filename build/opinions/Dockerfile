FROM golang:1.10.3 AS build
WORKDIR /go/src/github.com/chri5bot/opinions-pya/
COPY . .
RUN CGO_ENABLED=0 go install ./services/opinions

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /root/
COPY --from=build /go/bin/opinions .
ENTRYPOINT [ "./opinions" ]