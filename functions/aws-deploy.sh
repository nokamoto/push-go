#!/bin/bash

set -ex

cd functions/pub/aws

cp -r ../../google google

cp ../../../proto/service.proto .

# npm install

rm handler.zip

zip -r handler.zip index.js node_modules google service.proto config.json package.json
