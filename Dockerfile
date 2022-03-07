FROM golang:1.16-alpine

WORKDIR /app

COPY ./ app

COPY go.mod ./
COPY go.sum ./
RUN export GOPROXY=https://goproxy.io,direct
RUN go mod download

RUN go build -o /golang-echo-sample

EXPOSE 8080

CMD [ "/golang-echo-sample" ]