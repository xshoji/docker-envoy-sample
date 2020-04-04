FROM envoyproxy/envoy-dev:latest
RUN apt-get update
COPY ./config/envoy.yaml /etc/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy.yaml
