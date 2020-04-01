FROM gitpod/workspace-full@sha256:5c6442ee2382ec2f09a6c54568b9225ff4e19aa53876374fc65af0f3e42a2f05
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
