#!/usr/bin/env bash

dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
source "${dir}/../../helpers.bash"
# dir might have been overwritten by helpers.bash
dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

set -e

if [ -z "${K8S}" ] ; then
  log "K8S environment variable not set; please set it and re-run this script"
  exit 1
fi

case "${K8S}" in
  "1.6")
    NUM="6"
    ;;
  "1.7")
    NUM="7"
    ;;
  *)
    log "Usage: K8S={1.6,1.7} generate-certs.sh"
    exit 1
esac

export 'KUBERNETES_MASTER_IP4'=${KUBERNETES_MASTER_IP4:-"192.168.3$NUM.11"}
export 'KUBERNETES_MASTER_IP6'=${KUBERNETES_MASTER_IP6:-"FD01::B"}
export 'KUBERNETES_NODE_2_IP4'=${KUBERNETES_NODE_2_IP4:-"192.168.3$NUM.12"}
export 'KUBERNETES_NODE_2_IP6'=${KUBERNETES_NODE_2_IP6:-"FD01::C"}
export 'KUBERNETES_MASTER_SVC_IP4'=${KUBERNETES_MASTER_SVC_IP4:-"172.20.0.1"}
export 'KUBERNETES_MASTER_SVC_IP6'=${KUBERNETES_MASTER_SVC_IP6:-"FD03::1"}
export 'cluster_name'=${cluster_name:-"cilium-k8s-tests"}

log "KUBERNETES_MASTER_IP4: ${KUBERNETES_MASTER_IP4}"
log "KUBERNETES_MASTER_IP6: ${KUBERNETES_MASTER_IP6}"
log "KUBERNETES_NODE_2_IP4: ${KUBERNETES_NODE_2_IP4}"
log "KUBERNETES_NODE_2_IP6: ${KUBERNETES_NODE_2_IP6}"
log "KUBERNETES_MASTER_SVC_IP4: ${KUBERNETES_MASTER_SVC_IP4}"
log "KUBERNETES_MASTER_SVC_IP6: ${KUBERNETES_MASTER_SVC_IP6}"
log "cluster_name: ${cluster_name}"


function download_cfssl {
  wget --quiet https://pkg.cfssl.org/R1.2/cfssl_linux-amd64 > /usr/bin/cfssl && chmod +x /usr/bin/cfssl
  wget --quiet https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64 > /usr/bin/cfssljson && chmod +x /usr/bin/cfssljson
}

download_cfssl

if [ -z "$(command -v cfssl)" ]; then
    echo "cfssl not found, please download it from"
    echo "https://pkg.cfssl.org/R1.2/cfssl_linux-amd64"
    echo "and add it to your PATH."
    exit -1
fi

if [ -z "$(command -v cfssljson)" ]; then
    echo "cfssljson not found, please download it from"
    echo "https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64"
    echo "and add it to your PATH."
    exit -1
fi

log "creating ${dir}/ca-config.json"
cat > "${dir}/ca-config.json" <<EOF
{
  "signing": {
    "default": {
      "expiry": "8760h"
    },
    "profiles": {
      "kubernetes": {
        "usages": ["signing", "key encipherment", "server auth", "client auth"],
        "expiry": "8760h"
      }
    }
  }
}
EOF

log "creating ${dir}/ca-csr.json"
cat > "${dir}/ca-csr.json" <<EOF
{
  "CN": "Kubernetes",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "US",
      "L": "Portland",
      "O": "Kubernetes",
      "OU": "CA",
      "ST": "Oregon"
    }
  ]
}
EOF

log "generating certificates"
cfssl gencert -initca "${dir}/ca-csr.json" | cfssljson -bare "${dir}/ca"

log "creating ${dir}/kubernetes-csr.json"
cat > "${dir}/kubernetes-csr.json" <<EOF
{
  "CN": "kubernetes",
  "hosts": [
    "${KUBERNETES_MASTER_IP4}",
    "${KUBERNETES_MASTER_IP6}",
    "${KUBERNETES_MASTER_SVC_IP4}",
    "${KUBERNETES_MASTER_SVC_IP6}",
    "127.0.0.1",
    "::1",
    "localhost",
    "${cluster_name}.default"
  ],
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "US",
      "L": "Portland",
      "O": "Kubernetes",
      "OU": "Cluster",
      "ST": "Oregon"
    }
  ]
}
EOF

cfssl gencert \
  -ca="${dir}/ca.pem" \
  -ca-key="${dir}/ca-key.pem" \
  -config="${dir}/ca-config.json" \
  -profile=kubernetes \
  "${dir}/kubernetes-csr.json" | cfssljson -bare "${dir}/kubernetes"

rm "${dir}/ca-config.json" \
   "${dir}/ca-csr.json" \
   "${dir}/kubernetes-csr.json"
