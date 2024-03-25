FROM gitpod/workspace-full@sha256:4848c05c108e2f62ab14cae074a3528d1e6092d2ff003a6aa95992c756ec714a
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
