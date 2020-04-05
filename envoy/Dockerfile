FROM envoyproxy/envoy-dev:latest
RUN apt-get update
RUN apt-get install -y curl

CMD /usr/local/bin/envoy -c /opt/envoy/config/envoy.yaml
