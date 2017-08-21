#!/bin/bash

set -ex

cd docker

names=`echo push-go-{app,endpoint,log,notification,subscription}`

for name in $names
do
    sed -e "s/{NAME}/$name/g" Dockerfile.template > $name

    docker build -f $name -t nokamotohub/$name .

    docker push nokamotohub/$name
done

./test.sh

for name in $names
do
    docker push nokamotohub/$name
done
