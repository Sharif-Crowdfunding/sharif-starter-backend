
FROM golang:alpine

RUN go build -o build/server main.go

COPY ./build/server /bin/backend
CMD /bin/backend