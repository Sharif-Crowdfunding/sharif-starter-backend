
FROM golang
COPY ./build/server /bin/backend
CMD /bin/backend
#FROM alpine:3.13
#RUN apk --no-cache add ca-certificates
#WORKDIR /usr/bin
#COPY --from=build /go/src/app/bin /go/bin
#EXPOSE 80
#ENTRYPOINT /go/bin/web-app --port 80