FROM gitpod/workspace-full@sha256:5404ba92971a07c06855aa697a9cc490d190a110eab94c386c72f7ae20e53603
USER root
RUN echo "1" | apt-get update && apt-get install -y dh-make-golang
