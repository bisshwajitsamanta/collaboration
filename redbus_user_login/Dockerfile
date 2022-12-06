FROM golang:latest AS Builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o login cmd/server/main.go

FROM alpine:latest AS Production
COPY --from=Builder /app .
CMD ["./login"]
