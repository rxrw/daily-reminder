FROM golang:1.15.6-alpine

WORKDIR /app

ENV BASIC_USER user
ENV BASIC_PASS password
ENV DB_CHARSET utf8mb4
ENV LOTTERY_KEY 聚合彩票的key
ENV WEATHER_KEY 聚合天气的key
ENV TIANXING_KEY 天行数据api的key
ENV MYSQL_1_DBNAME dr
ENV MYSQL_1_HOST 127.0.0.1
ENV MYSQL_1_PASSWORD password
ENV MYSQL_1_PORT 3306
ENV MYSQL_1_USER user
ENV TZ Asia/Shanghai

ADD go.mod .

ADD go.sum .

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.io,direct

ADD . .

RUN go build main.go

EXPOSE 9000

CMD ["./main"]
