FROM alpine:3.12.0

COPY bin/build.sh .

RUN apk add --no-cache bash git && \
    bash build.sh && \
    rm build.sh && \
    apk del bash