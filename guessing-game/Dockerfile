FROM golang:1.14 AS builder

WORKDIR /run-app

COPY . /run-app

RUN go build -o gg


FROM scratch

WORKDIR /app

COPY --from=builder /run-app/gg /app/

ENTRYPOINT ["./gg"]