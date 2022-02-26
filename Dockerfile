FROM golang:1.17.5

# JSTをローカルタイムに
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

WORKDIR /go/src/