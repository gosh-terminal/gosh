FROM gitpod/workspace-full@sha256:00adf72da95e8f6ff4b8a59f90831ccf9a8d6b0ff8f63116f234f44e1d0a19cb
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
