FROM golang:stretch AS build

RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/install" | sh

FROM ubuntu

COPY --from=build /go/bin/gosh /go/bin
COPY --from=build /go/bin/history.txt /go/bin
CMD [ 'gosh', '-v' ]
