#!/bin/bash

set -ex

cd gcp

sed -e "s/{SERVICE_NAME}/$2/g" -e "s/{SERVICE_CONFIG_ID}/$3/g" k8s.$1.template.yaml > k8s.$1.yaml

kubectl create -f k8s.$1.yaml
