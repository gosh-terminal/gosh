FROM gitpod/workspace-full@sha256:5404ba92971a07c06855aa697a9cc490d190a110eab94c386c72f7ae20e53603
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/* \
USER gitpod
RUN curl -s "https://raw.githubusercontent.com/gosh-terminal/gosh/master/tools/setup2.0.bash" | bash \
    && chsh -s /home/gitpod/.gosh/gosh
