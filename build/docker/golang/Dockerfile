FROM golang:1.13.6-alpine

RUN apk add --update \
    curl \
    git \
    unzip \
    libzip-dev \
    libgcc \
    libbz2 \
    bzip2-dev \
    autoconf \
    make \
    zlib \
    zlib-dev \
    g++

RUN curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

RUN mkdir -p /var/billshare

COPY build/database/bills.json /var/billshare

WORKDIR /var/billshare

HEALTHCHECK \
  --interval=1m \
  --timeout=5s \
  CMD go version || exit 1