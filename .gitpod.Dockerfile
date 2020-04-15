FROM gitpod/workspace-full@sha256:881c732a9b82a99725620e9b5154620bbb7c0683e3cd9e21fbe246c14d2c4b67
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
