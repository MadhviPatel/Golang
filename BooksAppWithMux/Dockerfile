FROM golang:1.14

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/pmadhvi/Golang/BooksAppWithMux

COPY . ${SRC_DIR}

RUN cd ${SRC_DIR}; go build -o myapp; cp myapp /app/

ENTRYPOINT ["./myapp"]