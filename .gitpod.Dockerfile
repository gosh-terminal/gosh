FROM gitpod/workspace-full@sha256:a286eacd645a06c022f03a1edde2d8976877a146fd56e14ecf543418d0a6f3ab
USER root
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y \
    dh-make-golang --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*
