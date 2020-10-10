FROM golang:1.14

WORKDIR /app

COPY . /app

RUN go build -o guessing-game

ENTRYPOINT ["./guessing-game"]