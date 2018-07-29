FROM golang:alpine as server
ENV PKG=github.com/ivanturianytsia/raspi-temp
ENV DIR=$GOPATH/src/$PKG
RUN apk add --no-cache git && \
    go get \
    	github.com/gorilla/mux \
    	github.com/joho/godotenv \
      && \
    apk del git && \
    mkdir -p $DIR && \
    addgroup -S http && adduser -S -G http http
COPY ./server $DIR/server
WORKDIR $DIR
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -s' -installsuffix cgo -v -o /server $PKG/server

FROM alpine:3.7
LABEL maintainer="ivanturianytsia.work@gmail.com"
RUN apk --update upgrade && \
    apk --update add ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*
COPY --from=server /etc/passwd /etc/group /etc/
COPY --from=server /server /
RUN mkdir /data
RUN chmod 777 /data
USER http:http
EXPOSE 3000
ENTRYPOINT /server
