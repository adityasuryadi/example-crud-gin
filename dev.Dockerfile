FROM golang:1.19-alpine as builder

RUN apk update && apk add --no-cache git \
curl

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD ["air"]