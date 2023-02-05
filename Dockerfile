# syntax=docker/dockerfile:1

FROM golang:1.20-bullseye

WORKDIR /app

COPY *.go ./

RUN go build wtf.go

EXPOSE 8080:8080

CMD [ "./wtf" ]

