FROM golang:stretch AS build

RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/install" | sh

FROM ubuntu

COPY --from=build /go/bin/gosh /usr/bin/gosh
COPY --from=build /go/bin/history.txt $HOME/history.txt

ENV PATH=$HOME:$PATH

CMD [ 'gosh', '-v' ]
