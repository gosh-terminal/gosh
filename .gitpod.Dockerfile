FROM gitpod/workspace-full@sha256:702f40f1050af56dc37b0ab78c668976f6e8a28b778043665b1668350a617718
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
