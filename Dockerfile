# syntax=docker/dockerfile:1
# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM scratch

COPY --from=builder /app/.env.development /
COPY --from=builder /app/real-estate /
EXPOSE 8080
ENTRYPOINT ["./real-estate"]