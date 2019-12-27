FROM golang:stretch@sha256:0438f78156f40bde086dd13aaef23d8126e42d0a3074654d60fe115f36d9e81a

RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/setup2.0.bash" | bash

CMD [ "gosh", "-v" ]
