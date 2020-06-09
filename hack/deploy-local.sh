#!/usr/bin/env bash
# Deploy VMOP to the given cluster
#
# Usage:
# $ deploy-local.sh <deploy_yaml> <vmclasses_yaml>

set -o errexit
set -o pipefail
set -o nounset

YAML=$1
VMCLASSES_YAML=$2

KUBECTL="kubectl"

VMOP_NAMESPACE="vmware-system-vmop"
VMOP_DEPLOYMENT="vmware-system-vmop-controller-manager"

DEPLOYMENT_EXISTS=""
if $KUBECTL get deployment -n ${VMOP_NAMESPACE} ${VMOP_DEPLOYMENT} >/dev/null 2>&1 ; then
    DEPLOYMENT_EXISTS=1
fi

# Hack to reduce the number of replicas deployed from 3 to 1
# when deploying onto a single node kind cluster.
NODE_COUNT=$(kubectl get node --no-headers 2>/dev/null | wc -l)
if [ "$NODE_COUNT" -eq 1 ]; then
  sed -i -e 's/replicas: 3/replicas: 1/g' "$YAML"
  # remove the generated '-e' file on Mac
  rm -f "$YAML-e"
fi

$KUBECTL apply -f "$YAML"

if [[ -n $DEPLOYMENT_EXISTS ]] ; then
    $KUBECTL rollout restart -n ${VMOP_NAMESPACE} deployment ${VMOP_DEPLOYMENT}
    $KUBECTL rollout status -n ${VMOP_NAMESPACE} deployment ${VMOP_DEPLOYMENT}
fi

# Hack that retries applying the default VM Classes until the
# validating webhook is available.
VMOP_VMCLASSES_ATTEMPTS=0
while true ; do
    kubectl apply -f "${VMCLASSES_YAML}" && break
    VMOP_VMCLASSES_ATTEMPTS=$((VMOP_VMCLASSES_ATTEMPTS+1))
    if [[ $VMOP_VMCLASSES_ATTEMPTS -ge 60 ]] ; then
        echo "Failed to apply default VM Classes"
        exit 1
    fi
    echo "Cannot create default VM Classes. Trying again."
    sleep "5s"
done
