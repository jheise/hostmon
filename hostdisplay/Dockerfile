FROM jheise/ubuntu-golang

# RUN mkdir -p /go/src/hostagent
# ADD *.go /go/src/hostagent/
RUN go get github.com/jheise/hostmon/hostdisplay
RUN go install github.com/jheise/hostmon/hostdisplay
COPY  templates /srv
COPY static /srv
ENV ADDRESS 0.0.0.0
ENV PORT 9999
ENV REDISADDR redis
ENV REDISPORT 6379
ENV TEMPLATES /srv/templates
ENV STATIC /srv/static
EXPOSE ${PORT}
WORKDIR /srv
CMD /go/bin/hostdisplay -address ${ADDRESS} -port ${PORT} -redisaddr ${REDISADDR} -redisport ${REDISPORT}  -templates ${TEMPLATES} -static ${STATIC}
