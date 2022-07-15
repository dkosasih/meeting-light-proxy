FROM golang:1.18.4-alpine3.16 as builder

COPY . /go/src/github.com/dkosasih/meeting-light-proxy/
WORKDIR /go/src/github.com/dkosasih/meeting-light-proxy/
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/meeting-light-proxy .

FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/dkosasih/meeting-light-proxy/build/meeting-light-proxy /usr/bin/meeting-light/
COPY --from=builder /go/src/github.com/dkosasih/meeting-light-proxy/static/devcerts/localhost.key /usr/bin/meeting-light/static/devcerts/localhost.key
COPY --from=builder /go/src/github.com/dkosasih/meeting-light-proxy/static/devcerts/loclhost.crt /usr/bin/meeting-light/static/devcerts/loclhost.crt

WORKDIR /usr/bin/meeting-light/

ENTRYPOINT ["./meeting-light-proxy"]
