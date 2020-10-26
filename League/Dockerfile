FROM golang:1.14.3-alpine AS builder

WORKDIR /build

# Copies files and directories from host to image under /run-app/
COPY . .

RUN go build -o league .

WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/league .


FROM scratch AS bin

COPY --from=builder /dist/league /

ENTRYPOINT ["./league"]