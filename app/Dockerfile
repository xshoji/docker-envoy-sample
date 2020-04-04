# Base image
FROM alpine:3.11.3

# commands
RUN apk update \
  && apk upgrade \
  && apk add \
  bash \
  curl \
  less \
  vim \
  go \
;

# timezone
RUN apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

# copy app src
COPY ./app/src/main.go /tmp/main.go

ENTRYPOINT ["go", "run", "/tmp/main.go"]
