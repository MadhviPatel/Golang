FROM golang:1.14 AS builder

WORKDIR /app

COPY . /app

RUN GOOS=linux go build -o todoapp


FROM scratch

WORKDIR /run-app

COPY --from=builder /app/todoapp /run-app/

ENTRYPOINT ["./todoapp"]