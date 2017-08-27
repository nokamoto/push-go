#!/bin/bash

set -ex

cd gcp

protoc -I../proto --include_imports --include_source_info ../proto/service.proto --descriptor_set_out out.pb

sed -e "s/{PROJECT_ID}/$1/g" api_config.app.template.yaml > api_config.app.yaml
sed -e "s/{PROJECT_ID}/$1/g" api_config.endpoint.template.yaml > api_config.endpoint.yaml
sed -e "s/{PROJECT_ID}/$1/g" api_config.log.template.yaml > api_config.log.yaml
sed -e "s/{PROJECT_ID}/$1/g" api_config.notification.template.yaml > api_config.notification.yaml
sed -e "s/{PROJECT_ID}/$1/g" api_config.subscription.template.yaml > api_config.subscription.yaml

gcloud service-management deploy out.pb api_config.app.yaml
gcloud service-management deploy out.pb api_config.endpoint.yaml
gcloud service-management deploy out.pb api_config.log.yaml
gcloud service-management deploy out.pb api_config.notification.yaml
gcloud service-management deploy out.pb api_config.subscription.yaml
