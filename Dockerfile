FROM golang:stretch

RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/setup2.0.bash" | bash; . ~/.bashrc

CMD [ "gosh", "-v" ]
