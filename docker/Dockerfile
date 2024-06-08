FROM golang:1.22-bookworm

WORKDIR /app

RUN apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get -yq --no-install-recommends install \
  ca-certificates \
  sqlite3=3.* && \
  rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
  go install github.com/air-verse/air@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest

VOLUME /usr/src/app

CMD ["air", "-c", ".air.toml"]
EXPOSE 1323
