FROM gitpod/workspace-full@sha256:0e4ea1189b4c5172cbf0e5b74fe8703c71b3fa705edb12e8a7c2c5ccd81de986
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
