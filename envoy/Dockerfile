FROM envoyproxy/envoy-dev:latest
RUN apt-get update
RUN apt-get install -y curl
# nohup /usr/local/bin/envoy -c /opt/envoy/config/envoy2.yaml > /tmp/log.txt &

CMD ["/bin/bash"]
