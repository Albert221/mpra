FROM golang:1.14 as builder

WORKDIR /mpra
COPY . .
RUN go get github.com/markbates/pkger/cmd/pkger
RUN pkger
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mpra .

FROM alpine:latest

ENV HOST=
ENV PORT=8080
ENV REFRESH=1h

RUN apk --no-cache add ca-certificates
WORKDIR /mpra
COPY --from=builder /mpra .

CMD ./mpra -host=$HOST -port=$PORT -refresh=$REFRESH
