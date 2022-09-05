
#FROM golang:alpine
#RUN go mod init
#RUN go build -o build/server main.go
#
#COPY ./build/server /bin/backend
#CMD /bin/backend

FROM golang:alpine

WORKDIR /go/src/app

ADD . .
RUN go mod init

RUN go build  -o /server main.go

EXPOSE 8080

CMD ["./server"]