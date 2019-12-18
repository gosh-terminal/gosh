FROM golang:stretch AS build

RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/install" | sh

FROM ubuntu

COPY --from=build /go/bin/gosh /home/go/bin/gosh
COPY --from=build /go/bin/history.txt /home/go/bin/history.txt
ENV GOPATH=/home/go \
    PATH=$PATH:~$GOPATH/bin \
CMD [ 'gosh', '-v' ]
