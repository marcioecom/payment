FROM --platform=linux/amd64 golang:1.22

WORKDIR /go/src

RUN apt-get update && apt-get install gcc g++ libc-dev build-essential librdkafka-dev -y

CMD [ "tail", "-f", "/dev/null" ]
