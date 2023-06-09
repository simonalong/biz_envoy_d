#build stage
FROM golang:1.18 as builder


RUN go version

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app
EXPOSE 18003

#image stage
FROM envoyproxy/envoy:v1.21.2

WORKDIR /app

ENV TZ=Asia/Shanghai
ENV ZONEINFO=/app/zoneinfo.zip

COPY --from=builder /app/application.yaml /app/application.yaml
COPY --from=builder /app/biz-d-service /app/server
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /app

COPY ./envoy/start_service.sh /usr/local/bin/start_service.sh
COPY ./envoy/envoy-config.yaml /etc/envoy/envoy.yaml
ADD ./envoy/log /var/log/envoy

RUN chmod u+x /usr/local/bin/start_service.sh
ENTRYPOINT /usr/local/bin/start_service.sh

#CMD ["./server"]
