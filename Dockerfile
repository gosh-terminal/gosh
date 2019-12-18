FROM golang:stretch

RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/install" | sh


CMD [ "gosh", "-v" ]
