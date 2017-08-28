#!/bin/bash

set -ex

cd functions/pub/gcp

cp -r ../../google google

rm -rf node_modules
rm -f package-lock.json

gcloud beta functions deploy pub --stage-bucket $1 --trigger-http

cd ../../..
cd functions/resolv/gcp

cp -r ../../google google

rm -rf node_modules
rm -f package-lock.json

gcloud beta functions deploy resolv --stage-bucket $2 --trigger-topic resolve-endpoints
