#!/usr/bin/env bash

set -eo pipefail

DOCTL_CHECKSUM=1114bb864506c0371b469bb2edd4d544f6e902650cc259a273dab763956d94dd
DOCTL_VERSION=1.46.0

KUBECTL_CHECKSUM=1590d4357136a71a70172e32820c4a68430d1b94cf0ac941ea17695fbe0c5440d13e26e24a2e9ebdd360c231d4cd16ffffbbe5b577c898c78f7ebdc1d8d00fa3
KUBECTL_VERSION=1.19.0

setup() {
    mkdir /lib64
    ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
    apk add --no-cache wget
}

teardown() {
    apk del wget
}

install_doctl() {
    wget https://github.com/digitalocean/doctl/releases/download/v${DOCTL_VERSION}/doctl-${DOCTL_VERSION}-linux-amd64.tar.gz

    echo "${DOCTL_CHECKSUM}  doctl-${DOCTL_VERSION}-linux-amd64.tar.gz" > doctl.checksum
    sha256sum -c doctl.checksum
    rm doctl.checksum

    tar zxvf doctl-${DOCTL_VERSION}-linux-amd64.tar.gz
    mv doctl /usr/local/bin/doctl
    chmod +x /usr/local/bin/doctl

    rm -rf doctl-${DOCTL_VERSION}-linux-amd64.tar.gz
}



install_kubectl() {
    wget https://dl.k8s.io/v${KUBECTL_VERSION}/kubernetes-client-linux-amd64.tar.gz

    echo "${KUBECTL_CHECKSUM}  kubernetes-client-linux-amd64.tar.gz" > kubectl.checksum
    sha512sum -c kubectl.checksum
    rm kubectl.checksum
    
    tar zxvf kubernetes-client-linux-amd64.tar.gz
    mv kubernetes/client/bin/kubectl /usr/local/bin/kubectl
    chmod +x /usr/local/bin/kubectl
    
    rm -rf kubernetes-client-linux-amd64.tar.gz kubernetes
}


setup
install_kubectl
install_doctl
teardown