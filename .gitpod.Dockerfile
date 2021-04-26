FROM gitpod/workspace-full@sha256:42181fe6917aedc92d61d9fbcfbfe28feb5c4d7b5db27cf6817e7a7c2b9df573
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
