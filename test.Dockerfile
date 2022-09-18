# Builder
FROM golang:1.19-alpine as builder


RUN apk update && apk add --no-cache git \
    make openssh-client

WORKDIR /app
COPY . .
ENV MYSQL_TEST_URL=root:root@tcp(pos_mysql_test:3306)/inventory?parseTime=1
# RUN make full-test
# RUN cd /app/app \
# go test -v

CMD CGO_ENABLED=0 go test -v  ./...