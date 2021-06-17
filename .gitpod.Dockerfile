FROM gitpod/workspace-full@sha256:fc8753ebec4513625cab1d6ab3e61a2dbad02629a256f62525f91d971216073f
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
