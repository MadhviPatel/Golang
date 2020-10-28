FROM golang:1.14.3-alpine AS builder

WORKDIR /build

# Copies files and directories from host to image under /run-app/
COPY . .

RUN go build -o booksapp .

WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/booksapp .


FROM scratch AS bin

COPY --from=builder /dist/booksapp /

ENTRYPOINT ["./booksapp"]