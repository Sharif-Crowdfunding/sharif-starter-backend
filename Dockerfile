
#FROM golang:alpine
#RUN go mod init
#RUN go build -o build/server main.go
#
#COPY ./build/server /bin/backend
#CMD /bin/backend

FROM golang:alpine

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN rm go.mod
RUN go mod init sharif-starter-backend
RUN go mod tidy
RUN go build  -o main .

CMD ["/app/main"]