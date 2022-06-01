FROM golang AS builder
ENV GOPROXY https://goproxy.cn,direct
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
WORKDIR /app
COPY . /app
RUN go get -d -v ./...
RUN go build -o /output/server main.go

FROM alpine
ENV TZ Asia/Shanghai
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /output/server /app/server
COPY --from=builder /app/configs /app/configs
EXPOSE 8080
VOLUME /app/configs
CMD ["/app/server"]