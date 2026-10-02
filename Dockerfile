FROM golang:1.26-bookworm AS builder

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    gcc \
    zlib1g-dev \
    git \
    wget && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# glibc_compatibility.h နှင့် .c ကို Build Container ထဲ၌ တိုက်ရိုက် ဖန်တီးပေးခြင်း
RUN echo '#pragma once\n#ifdef __cplusplus\nextern "C" {\n#endif\n#ifdef __GLIBC__\n#if __GLIBC__ > 2 || (__GLIBC__ == 2 && __GLIBC_MINOR__ >= 28)\n#include <resolv.h>\n__attribute__((weak)) int __dn_expand(const unsigned char *msg, const unsigned char *eomorig, const unsigned char *comp_dn, char *exp_dn, int length) { return dn_expand(msg, eomorig, comp_dn, exp_dn, length); }\n__attribute__((weak)) int __res_nquery(res_state statp, const char *dname, int class, int type, unsigned char *answer, int anslen) { return res_nquery(statp, dname, class, type, answer, anslen); }\n#endif\n#endif\n#ifdef __cplusplus\n}\n#endif' > glibc_compatibility.h && \
    echo '#include "glibc_compatibility.h"' > glibc_compatibility.c

RUN go run github.com/AshokShau/gotdbot/scripts/tools
RUN go run setup_ntgcalls.go

RUN CGO_ENABLED=1 CGO_CFLAGS="-I/app -I." GOOS=linux go build -ldflags="-w -s" -o main .

FROM debian:12-slim AS runtime

RUN apt-get update && apt-get install -y --no-install-recommends \
    ffmpeg \
    wget \
    unzip \
    curl \
    lsb-release \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

RUN wget -O /usr/local/bin/yt-dlp \
    https://github.com/yt-dlp/yt-dlp-nightly-builds/releases/latest/download/yt-dlp_linux \
    && chmod +x /usr/local/bin/yt-dlp

RUN curl -fsSL https://deno.land/install.sh | sh \
    && export DENO_INSTALL="/opt/deno" \
    && export PATH="$DENO_INSTALL/bin:$PATH" \
    && mv /root/.deno /opt/deno \
    && ln -sf /opt/deno/bin/deno /usr/local/bin/deno

RUN groupadd -r app && useradd -r -g app -m -d /home/app app

ENV DENO_INSTALL="/opt/deno"
ENV PATH="${DENO_INSTALL}/bin:${PATH}"
ENV HOME="/home/app"

COPY --from=builder --chown=app:app /app/main /usr/local/bin/app
COPY --from=builder --chown=app:app /app/libtdjson.so* /home/app/

RUN chown -R app:app /opt/deno

USER app

WORKDIR /home/app
ENTRYPOINT ["app"]
