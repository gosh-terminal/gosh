FROM gitpod/workspace-full@sha256:25718cdee428df7bb6a4fc99793e73b51ab5387085dcebd66273ff5734c64e48
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
