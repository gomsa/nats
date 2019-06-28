FROM bigrocs/golang-gcc:1.12 as builder

WORKDIR /go/src/github.com/gomsa/nats
COPY . .

RUN GO111MODULE=off
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/service

FROM bigrocs/alpine:ca-data

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /go/src/github.com/gomsa/nats/bin/service /usr/local/bin/
CMD ["service"]