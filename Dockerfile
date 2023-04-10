FROM golang:1.19.5 as builder

ARG SERVICE

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /opt/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-w -s" -o /go/bin/service /opt/app/cmd/app

FROM scratch

COPY --from=builder /go/bin/service /go/bin/service

COPY ./web/dist/web /web/dist/web

CMD ["/go/bin/service"]