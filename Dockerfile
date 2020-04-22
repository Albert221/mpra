FROM golang:1.10 as builder

WORKDIR /go/src/github.com/Albert221/medicinal-products-registry-api

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

ENV MPR_ADDR=:80

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/Albert221/medicinal-products-registry-api/app .

CMD ["./app"]
