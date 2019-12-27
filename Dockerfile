FROM golang:stretch@sha256:3997ddbf0c313c5b629eccbdf14e098f4d3ad23cb1f2d9f8cb66707c9ee4ce79

RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/setup2.0.bash" | bash

CMD [ "gosh", "-v" ]
