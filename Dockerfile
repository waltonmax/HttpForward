FROM golang:latest

LABEL maintainer="waltonmax.github.io"

WORKDIR /app

COPY app .

RUN chmod -R 777 .

EXPOSE 6060

ENTRYPOINT ["./HttpForward"]

