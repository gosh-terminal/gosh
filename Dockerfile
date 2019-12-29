FROM golang:stretch@sha256:a3ce41aa0572b8809a5a4438adbcecb3dc0dd46276f9836ad3fbbe83277650a0

RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/setup2.0.bash" | bash

CMD [ "gosh", "-v" ]
