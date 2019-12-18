FROM golang:latest

ENV GOPATH=$HOME/go \
    PATH=$PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR /tmp/

COPY ./** /tmp/
RUN git clone https://github.com/gosh-terminal/gosh.git && cd gosh && ./tools/setup.sh

CMD [ "gosh", "-v" ]
