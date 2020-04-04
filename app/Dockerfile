# Base image
FROM alpine:3.11.3

# commands
RUN apk update \
  && apk upgrade \
  && apk add \
  bash \
  curl \
  less \
  tzdata \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apk del tzdata \
  && rm -rf /var/cache/apk/* \
;

# copy app binary
COPY ./app/src/app /opt/app

ENTRYPOINT ["/opt/app"]
