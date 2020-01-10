FROM golang:stretch@sha256:face5ed97d5d603cf2e36c3e91707c86f7394ea9d7364fd1e86a9d05f02c535d
RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/setup2.0.bash" | bash

CMD [ "gosh", "-v" ]
