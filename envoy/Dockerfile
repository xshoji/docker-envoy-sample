FROM envoyproxy/envoy-dev:latest
RUN apt-get update
RUN apt-get install -y curl
COPY ./envoy/config/envoy.yaml /etc/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy.yaml
