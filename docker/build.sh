#!/bin/bash

set -ex

cd docker

for name in push-go-{app,endpoint,log,notification}
do
    sed -e "s/{NAME}/$name/g" Dockerfile.template > $name

    docker build -f $name -t nokamotohub/$name .

    docker push nokamotohub/$name
done
