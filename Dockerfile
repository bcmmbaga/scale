FROM golang:1.16-alpine as builder

RUN apk add --no-cache git build-base && \
    apk add --no-cache upx

WORKDIR /scale

COPY docker .

RUN go get -d ./ && \
    CGO_ENABLED=0 GOOS=linux go build -o scale -a -ldflags="-s -w" -installsuffix cgo cmd/main.go && \
    upx  scale

FROM alpine
COPY --from=builder /scale/scale .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

ENTRYPOINT [ "./scale"]