FROM golang:1.22 as builder
WORKDIR /app
COPY ./ ./
RUN CGO_ENABLED=0 go build -o nocgo .

FROM alpine
WORKDIR /app
COPY --from=builder /app/nocgo .
