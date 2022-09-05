
FROM golang:alpine
RUN go mod init
RUN go build -o build/server main.go

COPY ./build/server /bin/backend
CMD /bin/backend