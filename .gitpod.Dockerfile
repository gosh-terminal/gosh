FROM gitpod/workspace-full@sha256:ee2f624ca47c33e50c6de33023c37f945ad1304efa869e9410adcf3b2319eb1b
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
