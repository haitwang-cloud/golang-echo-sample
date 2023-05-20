FROM golang:1.16-alpine

COPY ./ /app/
WORKDIR /app

RUN go mod download

RUN go build -o /golang-echo-sample

EXPOSE 8080

CMD [ "/golang-echo-sample" ]