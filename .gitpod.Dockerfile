FROM gitpod/workspace-full@sha256:9248cec8450a37620183e5ca4f8eb17abe03e575b1856d3ff29df18184c6d480
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
