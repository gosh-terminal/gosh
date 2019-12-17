FROM golang:latest

ENV GOPATH=$HOME/go \
    PATH=$PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR /tmp/

COPY ./** /tmp/

RUN ./tools/setup.sh

CMD [ "gosh", "-v" ]
