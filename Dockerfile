# syntax=docker/dockerfile:1
FROM golang:alpine3.16

WORKDIR /app

#RUN mkdir /root/.cache/rod/browser/chromium-1033860

COPY . .

COPY go.mod ./

RUN go get

RUN go mod tidy

RUN  CGO_ENABLED=0 GOOS=linux go build -a -o ./real-estate

#COPY ./real-estate /real-estate

CMD [ "./real-estate" ]