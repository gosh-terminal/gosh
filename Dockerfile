FROM golang:latest

ENV GOPATH=$HOME/go \
    PATH=$PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR /tmp/

COPY ./** /tmp/

RUN ./tools/install

CMD [ "gosh", "-v" ]
