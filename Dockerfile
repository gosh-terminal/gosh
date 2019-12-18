FROM ubuntu

RUN apt-get update \
    && apt-get -y install \
        golang \
        git

ENV HOME=/home \
    GOPATH=$HOME/go \
    PATH=$PATH:$GOROOT/bin:$GOPATH/bin

RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/install" | sh


CMD [ "gosh", "-v" ]
