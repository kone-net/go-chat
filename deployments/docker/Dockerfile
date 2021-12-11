FROM hub.c.163.com/public/centos:7.0
RUN [ "mkdir", "/usr/local/gochat" ]
RUN [ "mkdir", "-p", "/usr/local/gochat/web/static/file" ]
WORKDIR /usr/local/gochat
COPY ../../bin/chat /usr/local/gochat
COPY ../../config.toml /usr/local/gochat
CMD [ "/usr/local/gochat/chat" ]