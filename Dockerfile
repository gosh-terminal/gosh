FROM golang:latest

ENV GOPATH=$HOME/go \
    PATH=$PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR /tmp/

RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/setup2.0.bash" | bash

CMD [ "gosh", "-v" ]
