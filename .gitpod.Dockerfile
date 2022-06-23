FROM gitpod/workspace-full@sha256:c1c6267d94727a42158d953ab96ad0bcf8e2e3edfd87223555f441131ca8a97c
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
