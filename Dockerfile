FROM gliderlabs/alpine:3.2
MAINTAINER Langston Barrett <langston@aster.is> (@siddharthist)

# BUILD: sudo docker build -t whereismypi .
# RUN:   sudo docker run -it whereismypi

# If this file doesn't immedately work for you, please submit a Github issue:
# https://github.com/dlion/WhereIsMyPi/issues

# This docker container should run and then stop immediately when the program
# has exited.

RUN apk update && apk add go && rm -rf /var/cache/apk/*

WORKDIR /whereismypi
ENV PATH $PATH:/whereismypi
ADD . /whereismypi
RUN go build .

CMD ["./whereismypi"]
