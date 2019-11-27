FROM ubuntu
RUN apt-get update \
    && apt-get install -y \
    git \
    golang

ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin
RUN git clone https://github.com/JesterOrNot/gosh.git \
    && cd gosh \
    && bash setup.sh
CMD [ "gosh" ]
