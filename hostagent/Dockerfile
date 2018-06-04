FROM jheise/ubuntu-golang

RUN mkdir -p /go/src/hostagent
ADD *.go /go/src/hostagent/
RUN go get hostagent
RUN go install hostagent
ENV ADDRESS redis
ENV PORT 6379
ENV INTERVAL 15
WORKDIR /srv
CMD /go/bin/hostagent -address ${ADDRESS} -port ${PORT} -internval ${INTERVAL}
