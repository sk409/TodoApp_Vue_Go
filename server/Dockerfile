FROM golang:1.13

COPY entrypoint.sh /opt/entrypoint.sh

RUN go get github.com/jinzhu/gorm github.com/lib/pq \
    && chmod u+x /opt/entrypoint.sh

WORKDIR /go/github.com/sk409/todo

#CMD ["/opt/entrypoint.sh"]