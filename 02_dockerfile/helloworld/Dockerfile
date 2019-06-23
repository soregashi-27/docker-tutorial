FROM quay.io/cybozu/ubuntu:18.04

COPY hello.js /hello.js

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash - \
    && apt-get update \
    && apt-get install -y --no-install-recommends nodejs \
    && rm -rf /var/lib/apt/lists/*

EXPOSE 3000

CMD ["node", "hello"]