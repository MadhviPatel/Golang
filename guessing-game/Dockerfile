FROM golang:1.14.3-alpine AS builder

WORKDIR /run-app

# Copies files and directories from host to image under /run-app/
COPY . .

RUN go build -o gg


FROM scratch

COPY --from=builder /run-app/gg /

ENTRYPOINT ["./gg"]