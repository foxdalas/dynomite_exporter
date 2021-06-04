ARG ARCH="amd64"
ARG OS="linux"
FROM quay.io/prometheus/busybox-${OS}-${ARCH}:latest
LABEL maintainer="Maxim Pogozhy <foxdalass@gmail.com>"

ARG ARCH="amd64"
ARG OS="linux"
COPY .build/${OS}-${ARCH}/dynomite_exporter /bin/dynomite_exporter

USER       nobody
ENTRYPOINT ["/bin/dynomite_exporter"]
EXPOSE     9150
