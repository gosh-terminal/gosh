FROM gitpod/workspace-full@sha256:2356ec7bb0dd676b28d5f5d5d4a9f4c9e593100b187df9e1f9173c4295007fe9
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
